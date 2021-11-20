package domain

type User struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Notes []*Note `json:"notes"`
}
