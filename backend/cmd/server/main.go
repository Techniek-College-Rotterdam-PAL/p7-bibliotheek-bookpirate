package main

import (
	"math/rand"
	"server/internal/server"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	server.Run()
}
