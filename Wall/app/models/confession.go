package models

type Confession struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	Thumbon   uint   `json:"thumbon"`
	Thumbdown uint   `json:"thumbdown"`
	Hidename  uint   `json:"hidename"`
	Open      uint   `json:"open"`
	Poster    string `json:"poster"`
}
