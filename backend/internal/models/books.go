package models

type Book struct {
	Name     string `json:"title"`
	Isbn     string `json:"string,omitempty"`
	Author   string `json:"author,omitempty"`
	Language string `json:"language,omitempty"`
}

type SearchRequest struct {
	Book
}

type FeedRequest struct {
	Limit int `json:"list"`
}
