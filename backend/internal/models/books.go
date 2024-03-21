package models

type Book struct {
	SearchRequest
}

type SearchRequest struct {
	Name     string `json:"title"`
	Isbn     string `json:"string,omitempty"`
	Author   string `json:"author,omitempty"`
	Language string `json:"language,omitempty"`
}
