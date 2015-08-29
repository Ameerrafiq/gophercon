package structs

//structs to define json format for movieslist
type Movies struct {
	Url  string        `json:"url,omitempty"`
	Type string        `json:"type,omitempty"`
	Name string        `json:"name,omitempty"`
	Cast []CastingList `json:"cast,omitempty"`
}

type CastingList struct {
	Url  string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
	Role string `json:"role,omitempty"`
}
