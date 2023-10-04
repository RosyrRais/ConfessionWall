package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"Wall/config/database"
	"Wall/config/router"
)

func main() {
	//	log.Fatal("test\n")
	database.Init()
	r := gin.Default()
	router.Init(r)
	err := r.Run()

	if err != nil {
		log.Fatal("error in main.go", err)
	}
}
