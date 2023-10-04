package models

type Confession struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	Thumbon   uint   `json:"thumbon"`
	Thumbdown uint   `json:"thumbdown"`
}
