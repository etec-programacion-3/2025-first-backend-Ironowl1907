package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/etec-programacion-3/2025-first-backend-Ironowl1907/db"
	"github.com/etec-programacion-3/2025-first-backend-Ironowl1907/models"
	"github.com/gorilla/mux"
)

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	db.DB.Find(&books)
	json.NewEncoder(w).Encode(&books)

}
func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Book
	fmt.Println(params)
	db.DB.First(&book, params["id"])
	if book.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
	}
	json.NewEncoder(w).Encode(book)

}

func PostBookHandler(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	createdBook := db.DB.Create(&book)

	err := createdBook.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // code 400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&book)
}

func PutBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PutBookHandler"))
}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteBookHandler"))
}

func GetSearchBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetSearchBooksHandler"))
}
