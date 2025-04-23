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
		return
	}
	json.NewEncoder(w).Encode(&book)

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
	vars := mux.Vars(r)
	id := vars["id"]

	var existingBook models.Book
	if result := db.DB.First(&existingBook, id); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book not found!"))
		return
	}

	var updatedBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}

	updatedBook.ID = existingBook.ID

	db.DB.Save(&updatedBook)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedBook)
}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	params := mux.Vars(r)
	db.DB.First(&book, params["id"])

	if book.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User Not found"))
		return
	}

	db.DB.Delete(&book)
	w.WriteHeader(http.StatusOK)

}

func GetSearchBookHandler(w http.ResponseWriter, r *http.Request) {
	// Get name from URL parameters
	vars := mux.Vars(r)
	title := vars["titulo"]
	author := vars["autor"]
	category := vars["categoria"]

	// Query the user by name
	var book models.Book
	result := db.DB.Where("titulo = ? AND autor = ? AND categoria = ?", title, author, category).First(&book)
	if result.Error != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	// Return user data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
