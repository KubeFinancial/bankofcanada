package valet

type ApiErrorResponse struct {
	Message string `json:"message"`
	Docs    string `json:"docs"`
}

type Detail struct {
	Name        string `json:"name,omitempty"`
	Label       string `json:"label"`
	Description string `json:"description,omitempty"`
	Link        string `json:"link,omitempty"`
	Dimension   struct {
		Key  string `json:"key"`
		Name string `json:"name"`
	} `json:"dimension,omitempty"`
}

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
	Observations []map[string]any  `json:"observations,omitempty"`
}
