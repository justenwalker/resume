package data

// Company lists work experience at a given company
type Company struct {
	Name     string    `json:"name"`
	Website  string    `json:"website"`
	Title    string    `json:"title"`
	Start    Date      `json:"start"`
	End      Date      `json:"end"`
	Roles    []Role    `json:"roles"`
	Projects []Project `json:"projects"`
	Tags     []string  `json:"tags,omitempty"`
}

// Role describes a role performed within a company
type Role struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Highlights  []string `json:"highlights,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

// Project describes a notable project performed while at the company
type Project struct {
	Name        string
	Description string
}
