package valet

import (
	"fmt"
)

const baseURL = "https://www.bankofcanada.ca/valet"

// ListSeries fetches the list of all series.
func ListSeries() (map[string]Detail, error) {
	return fetchDetails("/lists/series/json")
}

// ListGroups fetches the list of all groups.
func ListGroups() (map[string]Detail, error) {
	return fetchDetails("/lists/groups/json")
}

// SeriesInfo fetches the details of a specific series.
func SeriesInfo(seriesName string) (Detail, error) {
	resp, err := API(fmt.Sprintf("%s/series/%s/json", baseURL, seriesName))
	if err != nil {
		return Detail{}, fmt.Errorf("fetching series info for %s: %w", seriesName, err)
	}
	return resp.SeriesDetails, nil
}

// GroupInfo fetches the details of a specific group.
func GroupInfo(groupName string) (GroupDetails, error) {
	resp, err := API(fmt.Sprintf("%s/groups/%s/json", baseURL, groupName))
	if err != nil {
		return GroupDetails{}, fmt.Errorf("fetching group info for %s: %w", groupName, err)
	}
	return resp.GroupDetails, nil
}

// SeriesObservations fetches observations for series.
func SeriesObservations(seriesNames string, options ...*ObservationOptions) ([]SeriesObservation, error) {
	return fetchObservationsWithCheck(seriesNames, "/observations/%s/json", "at least one series name", options...)
}

// GroupObservations fetches observations for a group.
func GroupObservations(groupName string, options ...*ObservationOptions) ([]SeriesObservation, error) {
	return fetchObservationsWithCheck(groupName, "/observations/group/%s/json", "group name", options...)
}
