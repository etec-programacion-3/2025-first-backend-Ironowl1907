package main

import (
	"net/http"

	"github.com/etec-programacion-3/2025-first-backend-Ironowl1907/db"
	routes "github.com/etec-programacion-3/2025-first-backend-Ironowl1907/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()
	g_Router := mux.NewRouter()

	g_Router.HandleFunc("/", routes.HomeHadler)

	http.ListenAndServe(":8080", g_Router)
}
