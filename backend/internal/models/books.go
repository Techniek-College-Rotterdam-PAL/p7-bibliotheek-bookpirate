package models

type Book struct {
	Stock    int    `json:"stock,omitempty"`
	Name     string `json:"title" validate:"required" binding:"required"`
	Isbn     string `json:"isbn,omitempty" validate:"required" binding:"required"`
	Author   string `json:"author,omitempty" validate:"required" binding:"required"`
	Language string `json:"language,omitempty"`
}

type SearchRequest struct {
	Name string `json:"title" validate:"required" binding:"required"`
}

type DeleteBook struct {
	Isbn string `json:"isbn" binding:"required"`
}
type InsertBook struct {
	Book
}

type Reservation struct {
	Id           uint32 `json:"id"`
	ReservedIsbn string `json:"isbn" binding:"required"`
	TimeStamp    int64  `json:"time_stamp"`
}

type FeedRequest struct {
	Limit int `json:"list"`
}
