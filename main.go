package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/marwenbhriz/ga-backend/controllers/apicontroller"
)

func main() {

	router := gin.Default()

	log.Println("Books API start listen on port 8081.")

	// apicontroller routes
	router.GET("/api", apicontroller.Index)

	router.Run(":8081")

}
