package data

// Institution containts details about a higher education instutition that was attended, and the degree achieved.
type Institution struct {
	Name    string   `json:"name"`
	Start   Date     `json:"start"`
	End     Date     `json:"end"`
	Website string   `json:"website"`
	Degree  string   `json:"degree"`
	GPA     string   `json:"gpa"`
	Honors  []string `json:"honors"`
}
