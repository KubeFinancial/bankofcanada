package valet

import (
	"fmt"
	"net/url"
	"time"
)

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

	if (opts.StartDate != "" && opts.EndDate == "") || (opts.EndDate != "" && opts.StartDate == "") {
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
