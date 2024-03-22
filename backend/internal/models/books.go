package models

type Book struct {
	Name     string `json:"title" binding:"required"`
	Isbn     string `json:"isbn,omitempty" binding:"required"`
	Author   string `json:"author,omitempty" binding:"required"`
	Language string `json:"language,omitempty"`
}

type SearchRequest struct {
	Book
}

type InsertBook struct {
	Book
}

type FeedRequest struct {
	Limit int `json:"list"`
}
