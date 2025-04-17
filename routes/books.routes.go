package routes

import "net/http"

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetBooksHandler"))
}
func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetBookHandler"))
}

func PostBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PostBookHandler"))
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
