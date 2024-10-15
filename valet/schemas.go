package valet

// ApiErrorResponse represents an error response from the Bank of Canada Valet API.
type ApiErrorResponse struct {
	Message string `json:"message"`
	Docs    string `json:"docs"`
}

// Detail represents detailed information about a series or group.
type Detail struct {
	Label       string `json:"label"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Link        string `json:"link,omitempty"`
	Dimension   struct {
		Key  string `json:"key"`
		Name string `json:"name"`
	} `json:"dimension,omitempty"`
}

// SeriesObservation represents a single observation for a series.
type SeriesObservation struct {
	Value string `json:"v"`
}

// Observation represents a collection of series observations for a specific date or quarter.
type Observation struct {
	Date    string                       `json:"d,omitempty"`
	Quarter string                       `json:"q,omitempty"`
	Series  map[string]SeriesObservation `json:"-"`
}

// ApiResponse represents the general structure of a response from the Bank of Canada Valet API.
type ApiResponse struct {
	Terms struct {
		URL string `json:"url"`
	} `json:"terms"`
	SeriesDetail  map[string]Detail `json:"seriesDetail,omitempty"`
	SeriesDetails Detail            `json:"seriesDetails,omitempty"`
	GroupDetail   Detail            `json:"groupDetail,omitempty"`
	GroupDetails  struct {
		Detail
		GroupSeries map[string]Detail `json:"groupSeries,omitempty"`
	} `json:"groupDetails,omitempty"`
	Series       map[string]Detail `json:"series,omitempty"`
	Groups       map[string]Detail `json:"groups,omitempty"`
	Observations []Observation     `json:"observations,omitempty"`
}
