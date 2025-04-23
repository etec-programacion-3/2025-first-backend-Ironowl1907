package main

import (
	"net/http"

	"github.com/etec-programacion-3/2025-first-backend-Ironowl1907/db"
	"github.com/etec-programacion-3/2025-first-backend-Ironowl1907/models"
	routes "github.com/etec-programacion-3/2025-first-backend-Ironowl1907/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()
	db.DB.AutoMigrate(models.Book{})

	g_Router := mux.NewRouter()

	g_Router.HandleFunc("/", routes.HomeHadler)
	g_Router.HandleFunc("/libros", routes.GetBooksHandler).Methods("GET")
	g_Router.HandleFunc("/libros/{id}", routes.GetBookHandler).Methods("GET")
	g_Router.HandleFunc("/libros", routes.PostBookHandler).Methods("POST")
	g_Router.HandleFunc("/libros/{id}", routes.PutBookHandler).Methods("PUT")
	g_Router.HandleFunc("/libros/{id}", routes.DeleteBookHandler).Methods("DELETE")
	g_Router.HandleFunc("/libros/buscar", routes.GetSearchBookHandler).Methods("GET")

	http.ListenAndServe(":8080", g_Router)
}
