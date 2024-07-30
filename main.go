package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rasyidridha547/simple-book-crud/models"
	"github.com/rasyidridha547/simple-book-crud/routes"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found")
	}
}

func main() {
	// models.ConnectDatabase()
	// defer models.CloseDatabase()

	r := mux.NewRouter()
	routes.BookRoutes(r)
	routes.Health(r)
	http.Handle("/", r)

	// graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-shutdown
		log.Println("Shutting down server....")

		// close DB connection
		if err := models.CloseDatabase(); err != nil {
			log.Fatal("Error Closing the Database: ", err)
		}

		log.Println("Server successfully shutdown")
		os.Exit(0)
	}()

	port := os.Getenv("PORT")

	log.Printf("Starting server on %v", port)
	ports := fmt.Sprintf(":%v", port)
	log.Fatal(http.ListenAndServe(ports, r))
}
