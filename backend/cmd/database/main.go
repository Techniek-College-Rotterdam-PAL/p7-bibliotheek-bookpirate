package main

import (
	"fmt"
	"server/internal/database"
)

func main() {
	if err := database.RunDriver(); err != nil {
		fmt.Println(err)
	}
}
