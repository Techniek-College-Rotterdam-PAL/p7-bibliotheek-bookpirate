package server

import (
	"crypto/sha1"
	"crypto/sha256"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	mu sync.Mutex

	sha256Hasher = sha256.New()
	sha1Hasher   = sha1.New()
)

const (
	PathAPI string = "/api/v1/"
)

var middleWares = map[string]func(*gin.Context){
	PathAPI + "register:POST":        Register,
	PathAPI + "login:POST":           Login,
	PathAPI + "fetch-books:POST":     FeedBooks,
	PathAPI + "add-book:POST":        AddBook,
	PathAPI + "search-books:POST":    SearchBooks,
	PathAPI + "remove:DELETE":        RemoveBook,
	PathAPI + "reserve-book:POST":    ReserveBook,
	PathAPI + "fetch-all-books:POST": FetchBooks,
}
