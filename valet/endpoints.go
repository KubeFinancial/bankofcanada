package valet

import (
	"fmt"
	"net/url"
	"time"
)

const baseURL = "https://www.bankofcanada.ca/valet"

// ListSeries fetches the list of all series.
func ListSeries() (map[string]Detail, error) {
	endpointURL := fmt.Sprintf("%s/lists/series/json", baseURL)
	resp, err := API(endpointURL)
	if err != nil {
		return nil, fmt.Errorf("fetching series list: %w", err)
	}
	return resp.Series, nil
}

// ListGroups fetches the list of all groups.
func ListGroups() (map[string]Detail, error) {
	endpointURL := fmt.Sprintf("%s/lists/groups/json", baseURL)
	resp, err := API(endpointURL)
	if err != nil {
		return nil, fmt.Errorf("fetching groups list: %w", err)
	}
	return resp.Groups, nil
}

// Series fetches the details of a specific series.
func Series(seriesName string) (Detail, error) {
	if seriesName == "" {
		return Detail{}, fmt.Errorf("series name is required")
	}

	endpointURL := fmt.Sprintf("%s/observations/%s/json?recent=1", baseURL, seriesName)
	resp, err := API(endpointURL)
	if err != nil {
		return Detail{}, fmt.Errorf("fetching series list: %w", err)
	}

	return resp.SeriesDetail[seriesName], nil
}

// Group fetches the details of a specific group.
func Group(groupName string) (GroupDetails, error) {
	if groupName == "" {
		return GroupDetails{}, fmt.Errorf("group name is required")
	}

	endpointURL := fmt.Sprintf("%s/observations/group/%s/json?recent=1", baseURL, groupName)
	resp, err := API(endpointURL)
	if err != nil {
		return GroupDetails{}, fmt.Errorf("fetching group details: %w", err)
	}

	groupDetails := GroupDetails{
		Detail:      resp.GroupDetail,
		GroupSeries: resp.SeriesDetail,
	}

	// Set the Name field explicitly
	groupDetails.Detail.Name = groupName

	return groupDetails, nil
}

// SeriesObservations fetches observations for series.
func SeriesObservations(
	seriesNames string,
	options ...*ObservationOptions,
) ([]SeriesObservation, error) {
	if seriesNames == "" {
		return nil, fmt.Errorf("atleast one series name is required")
	}
	endpointURL := fmt.Sprintf("%s/observations/%s/json", baseURL, seriesNames)
	var opts *ObservationOptions
	if len(options) > 0 {
		opts = options[0]
	}
	return fetchObservations(endpointURL, opts)
}

// GroupObservations fetches observations for a group.
func GroupObservations(
	groupName string,
	options ...*ObservationOptions,
) ([]SeriesObservation, error) {
	if groupName == "" {
		return nil, fmt.Errorf("group name is required")
	}
	endpointURL := fmt.Sprintf("%s/observations/group/%s/json", baseURL, groupName)
	var opts *ObservationOptions
	if len(options) > 0 {
		opts = options[0]
	}
	return fetchObservations(endpointURL, opts)
}

func fetchObservations(endpointURL string, opts *ObservationOptions) ([]SeriesObservation, error) {
	u, err := url.Parse(endpointURL)
	if err != nil {
		return nil, fmt.Errorf("parsing URL: %w", err)
	}

	query := u.Query()
	if opts == nil || !(opts.Recent != 0 || opts.StartDate != "" || opts.EndDate != "" ||
		opts.RecentWeeks != 0 || opts.RecentMonths != 0 || opts.RecentYears != 0) {
		opts = &ObservationOptions{Recent: 1}
	}

	if (opts.StartDate != "" && opts.EndDate == "") ||
		(opts.EndDate != "" && opts.StartDate == "") {
		return nil, fmt.Errorf("both StartDate and EndDate must be provided")
	}

	if opts.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", opts.StartDate)
		if err != nil {
			return nil, fmt.Errorf("invalid StartDate: %w", err)
		}
		endDate, err := time.Parse("2006-01-02", opts.EndDate)
		if err != nil {
			return nil, fmt.Errorf("invalid EndDate: %w", err)
		}
		if endDate.Before(startDate) {
			return nil, fmt.Errorf("EndDate must be after StartDate")
		}
		query.Set("start_date", opts.StartDate)
		query.Set("end_date", opts.EndDate)
	} else {
		count := 0
		for key, value := range map[string]int{
			"recent":        opts.Recent,
			"recent_weeks":  opts.RecentWeeks,
			"recent_months": opts.RecentMonths,
			"recent_years":  opts.RecentYears,
		} {
			if value > 0 {
				query.Set(key, fmt.Sprintf("%d", value))
				count++
			}
		}
		if count > 1 {
			return nil, fmt.Errorf("only one time range option can be provided")
		}
	}

	if opts.OrderDir == "asc" || opts.OrderDir == "desc" {
		query.Set("order_dir", opts.OrderDir)
	}
	u.RawQuery = query.Encode()

	resp, err := API(u.String())
	if err != nil {
		return nil, fmt.Errorf("fetching observations: %w", err)
	}

	var result []SeriesObservation
	for _, obs := range resp.Observations {
		for name, seriesObs := range obs.Series {
			result = append(result, SeriesObservation{
				Date:    obs.Date,
				Quarter: obs.Quarter,
				Name:    name,
				Value:   seriesObs.Value,
			})
		}
	}
	return result, nil
}
