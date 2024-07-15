package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/marwenbhriz/ga-backend/controllers/apicontroller"
	"github.com/marwenbhriz/ga-backend/controllers/bookcontroller"
	"github.com/marwenbhriz/ga-backend/models"
)

func main() {

	router := gin.Default()

	models.ConnectDatabase()

	log.Println("Books API start listen on port 8091.")

	// apicontroller routes
	router.GET("/api", apicontroller.Index)

	// books routes
	router.GET("/api/books", bookcontroller.Index)
	router.GET("/api/book/:id", bookcontroller.Show)
	router.POST("/api/book", bookcontroller.Create)
	router.PUT("/api/book/:id", bookcontroller.Update)
	router.DELETE("/api/book", bookcontroller.Delete)

	router.Run(":8091")

}
