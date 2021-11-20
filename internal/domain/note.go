package domain

type Note struct {
	ID     string `json:"id"`
	UserID string `json:"-"`
	Text   string `json:"text"`
}
