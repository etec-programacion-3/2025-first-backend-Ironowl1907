package routes

import (
	"encoding/json"
	"net/http"

	"github.com/etec-programacion-3/2025-first-backend-Ironowl1907/db"
	"github.com/etec-programacion-3/2025-first-backend-Ironowl1907/models"
)

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetBooksHandler"))
}
func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetBookHandler"))
}

func PostBookHandler(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	createdUser := db.DB.Create(&book)

	err := createdUser.Error
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
