package data

// Writing contains information about a blog post or article I've authored.
type Writing struct {
	Title           string `json:"event"`
	Link            string `json:"link"`
	Date            Date   `json:"date"`
	Description     string `json:"description"`
	Publication     string `json:"publication"`
	PublicationLink string `json:"publication_link"`
}
