package valet

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
