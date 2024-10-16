package valet

import (
	"encoding/json"
)

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
