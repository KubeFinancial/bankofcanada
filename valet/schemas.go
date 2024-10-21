package valet

import (
	"encoding/json"
)

// SeriesObservation represents a single observation for a series.
type SeriesObservation struct {
	Date    string `json:"-"`
	Quarter string `json:"-"`
	Name    string `json:"-"`
	Value   string `json:"v,omitempty"`
}

// Observation represents a collection of series observations for a specific date or quarter.
type Observation struct {
	Date    string                       `json:"d,omitempty"`
	Quarter string                       `json:"q,omitempty"`
	Series  map[string]SeriesObservation `json:"-"`
}

// Detail represents detailed information about a series or group.
type Detail struct {
	Label       string `json:"label,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Link        string `json:"link,omitempty"`
	Dimension   struct {
		Key  string `json:"key"`
		Name string `json:"name"`
	} `json:"dimension,omitempty"`
}

// GroupDetails represents the structure of group details in the API response
type GroupDetails struct {
	Detail
	GroupSeries map[string]Detail `json:"groupSeries,omitempty"`
}

// APIResponse represents the general structure of a response from the Bank of Canada Valet API.
type APIResponse struct {
	// terms
	Terms struct {
		URL string `json:"url,omitempty"`
	} `json:"terms,omitempty"`

	// /lists/series/json
	Series map[string]Detail `json:"series,omitempty"`

	// /lists/groups/json
	Groups map[string]Detail `json:"groups,omitempty"`

	// /series/{seriesName}/json
	SeriesDetails Detail `json:"seriesDetails,omitempty"`

	// /groups/{groupName}/json
	GroupDetails GroupDetails `json:"groupDetails,omitempty"`

	// /observations/{seriesNames}/json
	// /observations/group/{groupName}/json
	SeriesDetail map[string]Detail `json:"seriesDetail,omitempty"`
	GroupDetail  Detail            `json:"groupDetail,omitempty"`
	Observations []Observation     `json:"observations,omitempty"`

	// 	Errors
	Message string `json:"message,omitempty"`
	Docs    string `json:"docs,omitempty"`
}

// ObservationOptions defines available query options for the API.
type ObservationOptions struct {
	StartDate    string
	EndDate      string
	Recent       int
	RecentWeeks  int
	RecentMonths int
	RecentYears  int
	OrderDir     string
}

// UnmarshalJSON Custom unmarshal function to populate the Name field in Detail
func (r *APIResponse) UnmarshalJSON(data []byte) error {
	type Alias APIResponse
	aux := struct{ *Alias }{Alias: (*Alias)(r)}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Populate Name for all Detail types
	populateNames := func(m map[string]Detail) {
		if m == nil {
			return
		}
		for key, detail := range m {
			detail.Name = key
			m[key] = detail
		}
	}

	populateNames(r.SeriesDetail)
	populateNames(r.GroupDetails.GroupSeries)
	populateNames(r.Series)
	populateNames(r.Groups)

	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for custom unmarshalling of Observation.
func (o *Observation) UnmarshalJSON(data []byte) error {
	raw := make(map[string]json.RawMessage)

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Helper function to unmarshal values
	unmarshalField := func(field string, target *string) error {
		if val, ok := raw[field]; ok {
			return json.Unmarshal(val, target)
		}
		return nil
	}

	// Unmarshal date and quarter using the helper function
	if err := unmarshalField("d", &o.Date); err != nil {
		return err
	}
	if err := unmarshalField("q", &o.Quarter); err != nil {
		return err
	}

	// Initialize the Series map
	o.Series = make(map[string]SeriesObservation)

	// Unmarshal series observations and set their names
	for key, value := range raw {
		if key == "d" || key == "q" {
			continue
		}

		var seriesObservation SeriesObservation
		if err := json.Unmarshal(value, &seriesObservation); err != nil {
			return err
		}
		seriesObservation.Name = key
		seriesObservation.Date = o.Date
		seriesObservation.Quarter = o.Quarter
		o.Series[key] = seriesObservation
	}

	return nil
}
