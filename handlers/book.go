package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rasyidridha547/simple-book-crud/models"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	models.DB.Find(&books)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Data:    books,
	})
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var book models.Book

	if err := models.DB.First(&book, params["name"]).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Error:   "Book Not Found!",
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Data:    book,
	})
}
