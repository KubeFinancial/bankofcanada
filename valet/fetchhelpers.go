package valet

import (
	"fmt"
	"net/url"
	"time"
)

// fetchDetails handles common logic for fetching lists
func fetchDetails(path string) (map[string]Detail, error) {
	resp, err := API(baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("fetching %s: %w", path, err)
	}
	if path == "/lists/series/json" {
		return resp.Series, nil
	}
	return resp.Groups, nil
}

func fetchObservationsWithCheck(name, pathFormat, errContext string, opts ...*ObservationOptions) ([]SeriesObservation, error) {
	if name == "" {
		return nil, fmt.Errorf("%s is required", errContext)
	}
	return fetchObservations(fmt.Sprintf(pathFormat, name), firstOption(opts))
}

func fetchObservations(path string, opts *ObservationOptions) ([]SeriesObservation, error) {
	URL, err := buildObservationURL(path, opts)
	if err != nil {
		return nil, fmt.Errorf("building observation URL: %w", err)
	}
	resp, err := API(URL)
	if err != nil {
		return nil, fmt.Errorf("fetching observations: %w", err)
	}
	return flattenObservations(resp.Observations), nil
}

func buildObservationURL(path string, opts *ObservationOptions) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("parsing base URL: %w", err)
	}
	u.Path += path

	query := u.Query()
	opts = setDefaultOptions(opts)
	if err := setTimeRangeParams(&query, opts); err != nil {
		return "", err
	}
	setOrderDirParam(&query, opts)
	u.RawQuery = query.Encode()

	return u.String(), nil
}

func setTimeRangeParams(q *url.Values, opts *ObservationOptions) error {
	if err := validateTimeRange(opts); err != nil {
		return err
	}
	if opts.StartDate != "" {
		q.Set("start_date", opts.StartDate)
		q.Set("end_date", opts.EndDate)
	}
	for key, value := range map[string]int{
		"recent":        opts.Recent,
		"recent_weeks":  opts.RecentWeeks,
		"recent_months": opts.RecentMonths,
		"recent_years":  opts.RecentYears,
	} {
		if value > 0 {
			q.Set(key, fmt.Sprintf("%d", value))
		}
	}
	return nil
}

func setOrderDirParam(q *url.Values, opts *ObservationOptions) {
	if opts.OrderDir == "asc" || opts.OrderDir == "desc" {
		q.Set("order_dir", opts.OrderDir)
	}
}

func validateTimeRange(opts *ObservationOptions) error {
	if (opts.StartDate != "" && opts.EndDate == "") || (opts.EndDate != "" && opts.StartDate == "") {
		return fmt.Errorf("both StartDate and EndDate must be provided")
	}
	if opts.StartDate != "" {
		return validateDateRange(opts.StartDate, opts.EndDate)
	}
	if countNonZero(opts.Recent, opts.RecentWeeks, opts.RecentMonths, opts.RecentYears) > 1 {
		return fmt.Errorf("only one time range option can be provided")
	}
	return nil
}

func validateDateRange(start, end string) error {
	startDate, err := time.Parse("2006-01-02", start)
	if err != nil {
		return fmt.Errorf("invalid StartDate: %w", err)
	}
	endDate, err := time.Parse("2006-01-02", end)
	if err != nil {
		return fmt.Errorf("invalid EndDate: %w", err)
	}
	if endDate.Before(startDate) {
		return fmt.Errorf("EndDate must be after StartDate")
	}
	return nil
}

func flattenObservations(observations []Observation) []SeriesObservation {
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

// Utility functions
func firstOption(opts []*ObservationOptions) *ObservationOptions {
	if len(opts) > 0 {
		return opts[0]
	}
	return nil
}

func setDefaultOptions(opts *ObservationOptions) *ObservationOptions {
	if opts == nil || !hasTimeOption(opts) {
		return &ObservationOptions{Recent: 1}
	}
	return opts
}

func hasTimeOption(opts *ObservationOptions) bool {
	return opts.Recent != 0 || opts.StartDate != "" || opts.EndDate != "" ||
		opts.RecentWeeks != 0 || opts.RecentMonths != 0 || opts.RecentYears != 0
}

func countNonZero(vals ...int) int {
	count := 0
	for _, v := range vals {
		if v > 0 {
			count++
		}
	}
	return count
}
