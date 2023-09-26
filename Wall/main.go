package main

import (
	"log"

	"Wall/config/database"

	"github.com/gin-gonic/gin"
)

func main() {
	//	log.Fatal("test\n")
	database.Init()
	r := gin.Default()

	err := r.Run()

	if err != nil {
		log.Fatal("error in main.go", err)
	}
}
