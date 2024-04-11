package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"server/internal/server"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	server.Run(gin.New())

}
