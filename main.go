package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
	"github.com/marwenbhriz/ga-backend/controllers/apicontroller"
	"github.com/marwenbhriz/ga-backend/controllers/usercontroller"
	"github.com/marwenbhriz/ga-backend/models"
)

func main() {

	router := gin.Default()

	models.ConnectDatabase()

	log.Println("Tasks API start listen on port 8086.")

	// apicontroller routes
	router.GET("/api", apicontroller.Index)

	// tasks routes
	router.GET("/api/users", usercontroller.Index)
	router.GET("/api/user/:id", usercontroller.Show)
	router.POST("/api/user", usercontroller.Create)
	router.PUT("/api/user/:id", usercontroller.Update)
	router.DELETE("/api/user", usercontroller.Delete)

	// cors middleware
	methods := handlers.AllowedMethods([]string{"OPTIONS", "DELETE", "GET", "HEAD", "POST", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	handler := handlers.CORS(methods, origins, headers)(router)

	// create and start server
	s := &http.Server{
		Addr:     ":8086",       // bind address
		Handler:  handler,       // default handler
		ErrorLog: log.Default(), // logger for the server
	}
	go func() {
		log.Println("Listening on port 8086")
		log.Fatal(s.ListenAndServe())
	}()

	// trap terminate or interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	// block shutdown until terminal/interrupt signal received
	// ⛔️ DON'T DELETE! This will shutdown server immediately
	sig := <-c
	log.Println("Gracefully shutting down...", sig)

	// gracefully shutdown server, waiting max 30 seconds for current operations to complete
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	cancel()
	s.Shutdown(tc)

	//router.Run(":8096")

}
