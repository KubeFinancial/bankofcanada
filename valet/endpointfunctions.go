package valet

import (
	"fmt"
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

// SeriesInfo fetches the details of a specific series.
func SeriesInfo(seriesName string) (Detail, error) {
	if seriesName == "" {
		return Detail{}, fmt.Errorf("series name is required")
	}
	endpointURL := fmt.Sprintf("%s/series/%s/json", baseURL, seriesName)
	resp, err := API(endpointURL)
	if err != nil {
		return Detail{}, fmt.Errorf("fetching series info for %s: %w", seriesName, err)
	}
	return resp.SeriesDetails, nil
}

// GroupInfo fetches the details of a specific group.
func GroupInfo(groupName string) (GroupDetails, error) {
	if groupName == "" {
		return GroupDetails{}, fmt.Errorf("group name is required")
	}
	endpointURL := fmt.Sprintf("%s/groups/%s/json", baseURL, groupName)
	resp, err := API(endpointURL)
	if err != nil {
		return GroupDetails{}, fmt.Errorf("fetching group info for %s: %w", groupName, err)
	}
	return resp.GroupDetails, nil
}

// SeriesObservations fetches observations for series.
func SeriesObservations(seriesNames string, options ...*ObservationOptions) ([]SeriesObservation, error) {
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
func GroupObservations(groupName string, options ...*ObservationOptions) ([]SeriesObservation, error) {
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

func AwesomeSeries(seriesName string) (Detail, error) {
	if seriesName == "" {
		return Detail{}, fmt.Errorf("series name is required")
	}
	endpointURL := fmt.Sprintf("%s/observations/%s/json?recent=1", baseURL, seriesName)
	resp, err := API(endpointURL)
	if err != nil {
		return Detail{}, fmt.Errorf("fetching series list: %w", err)
	}
	var result Detail
	for _, series := range resp.SeriesDetail {
		result = series
	}
	return result, nil
}
