package valet

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

const (
	baseURL = "https://www.bankofcanada.ca/valet"
)

// ListSeries fetches the list of all series.
func ListSeries() (map[string]Detail, error) {
	endpointURL := baseURL + "/lists/series/json"

	resp, err := API(endpointURL)
	if err != nil {
		return nil, fmt.Errorf("fetching series list: %w", err)
	}

	return resp.Series, nil
}

// ListGroups fetches the list of all groups.
func ListGroups() (map[string]Detail, error) {
	endpointURL := baseURL + "/lists/groups/json"

	resp, err := API(endpointURL)
	if err != nil {
		return nil, fmt.Errorf("fetching groups list: %w", err)
	}

	return resp.Groups, nil
}

// Series fetches the details of a specific series.
func Series(seriesName string) (Detail, error) {
	if seriesName == "" {
		return Detail{}, errors.New("series name is required")
	}

	resp, err := fetchObservations(seriesName, false, nil)
	if err != nil {
		return Detail{}, fmt.Errorf("fetching series details: %w", err)
	}

	return resp.SeriesDetail[seriesName], nil
}

// Group fetches the details of a specific group.
func Group(groupName string) (GroupDetails, error) {
	if groupName == "" {
		return GroupDetails{}, errors.New("group name is required")
	}

	resp, err := fetchObservations(groupName, true, nil)
	if err != nil {
		return GroupDetails{}, fmt.Errorf("fetching group details: %w", err)
	}

	groupDetails := GroupDetails{
		Detail:      resp.GroupDetail,
		GroupSeries: resp.SeriesDetail,
	}
	groupDetails.Detail.Name = groupName

	return groupDetails, nil
}

// SeriesObservations fetches observations for series.
func SeriesObservations(
	seriesNames string,
	options ...*ObservationOptions,
) ([]SeriesObservation, error) {
	if seriesNames == "" {
		return nil, errors.New("at least one series name is required")
	}

	resp, err := fetchObservations(seriesNames, false, options...)
	if err != nil {
		return nil, err
	}

	return parseObservations(resp.Observations), nil
}

// GroupObservations fetches observations for a group.
func GroupObservations(
	groupName string,
	options ...*ObservationOptions,
) ([]SeriesObservation, error) {
	if groupName == "" {
		return nil, errors.New("group name is required")
	}

	resp, err := fetchObservations(groupName, true, options...)
	if err != nil {
		return nil, err
	}

	return parseObservations(resp.Observations), nil
}

// fetchObservations is an internal function that fetches data from the API.
func fetchObservations(
	name string,
	isGroup bool,
	options ...*ObservationOptions,
) (APIResponse, error) {
	var endpointURL string
	if isGroup {
		endpointURL = baseURL + "/observations/group/" + name + "/json"
	} else {
		endpointURL = baseURL + "/observations/" + name + "/json"
	}

	parsedURL, err := url.Parse(endpointURL)
	if err != nil {
		return APIResponse{}, fmt.Errorf("parsing URL: %w", err)
	}

	query := parsedURL.Query()

	var opts *ObservationOptions
	if len(options) > 0 && options[0] != nil {
		opts = options[0]
	} else {
		opts = &ObservationOptions{
			Recent:       1,
			StartDate:    "",
			EndDate:      "",
			RecentWeeks:  0,
			RecentMonths: 0,
			RecentYears:  0,
			OrderDir:     "",
		}
	}

	err = applyOptions(opts, query)
	if err != nil {
		return APIResponse{}, err
	}

	parsedURL.RawQuery = query.Encode()

	return API(parsedURL.String())
}

func applyOptions(opts *ObservationOptions, query url.Values) error {
	if opts.StartDate != "" || opts.EndDate != "" {
		if opts.StartDate == "" || opts.EndDate == "" {
			return errors.New("both StartDate and EndDate must be provided")
		}

		startDate, err := time.Parse("2006-01-02", opts.StartDate)
		if err != nil {
			return fmt.Errorf("invalid StartDate: %w", err)
		}

		endDate, err := time.Parse("2006-01-02", opts.EndDate)
		if err != nil {
			return fmt.Errorf("invalid EndDate: %w", err)
		}

		if endDate.Before(startDate) {
			return errors.New("EndDate must be after StartDate")
		}

		query.Set("start_date", opts.StartDate)
		query.Set("end_date", opts.EndDate)
	} else {
		setIfPositive := func(key string, value int) {
			if value > 0 {
				query.Set(key, strconv.Itoa(value))
			}
		}
		setIfPositive("recent", opts.Recent)
		setIfPositive("recent_weeks", opts.RecentWeeks)
		setIfPositive("recent_months", opts.RecentMonths)
		setIfPositive("recent_years", opts.RecentYears)
	}

	if opts.OrderDir == "asc" || opts.OrderDir == "desc" {
		query.Set("order_dir", opts.OrderDir)
	}

	return nil
}

func parseObservations(observations []Observation) []SeriesObservation {
	var result []SeriesObservation

	for _, obs := range observations {
		for name, seriesObs := range obs.Series {
			result = append(result, SeriesObservation{
				Date:    obs.Date,
				Quarter: obs.Quarter,
				Name:    name,
				Value:   seriesObs.Value,
			})
		}
	}

	return result
}
