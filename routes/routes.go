package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rasyidridha547/simple-book-crud/handlers"
)

func BookRoutes(router *mux.Router) {
	router.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{name}", handlers.GetBook).Methods("GET")
}

func Health(router *mux.Router) {
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Service is up!"))
	})
}
