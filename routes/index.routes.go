package routes

import "net/http"

func HomeHadler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}
