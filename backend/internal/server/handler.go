package server

import "github.com/gin-gonic/gin"

const (
	PathAPI string = "/api/v1/"
)

var middleWares = map[string]func(*gin.Context){
	PathAPI + "register:POST":     Register,
	PathAPI + "fetch-books:POST":  FeedBooks,
	PathAPI + "search-books:POST": nil,
	PathAPI + "remove:DELETE":     nil,
}
