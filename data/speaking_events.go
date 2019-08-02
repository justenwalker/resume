package data

// SpeakingEvent contains information about a talk or presentation that I gave at a conference or event
type SpeakingEvent struct {
	Name        string `json:"name"`
	Event       string `json:"event"`
	Date        Date   `json:"date"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Notes       string `json:"notes,omitempty"`
}
