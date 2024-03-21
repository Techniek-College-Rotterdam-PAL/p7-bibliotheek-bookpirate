package server

import "github.com/gin-gonic/gin"

const (
	PathAPI string = "/api/v1/"
)

var middleWares = map[string]func(*gin.Context){
	PathAPI + "register:POST": Register,
	PathAPI + "search:POST":   nil,
	PathAPI + "remove:DELETE": nil,
}
