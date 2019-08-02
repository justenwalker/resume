package data

// OpenSource contains details about Open Source projects that have been created or contributed to
type OpenSource struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Website      string   `json:"website"`
	Link         string   `json:"link,omitempty"`
	Contribution string   `json:"contribution"`
	Tags         []string `json:"tags,omitempty"`
}

const (
	// MinorContributor indicates small bugs or patches were provided
	MinorContributor = "minor"
	// MajorContributor indicates large contributions were made
	MajorContributor = "major"
	// Creator means the project was created by me
	Creator = "creator"
)
