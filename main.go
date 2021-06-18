package main

import (
	"mux-http-api/app"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/posts", app.AddPost).Methods("POST")
	router.HandleFunc("/posts", app.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", app.GetPost).Methods("GET")
	router.HandleFunc("/posts/{id}", app.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", app.DeletePost).Methods("DELETE")
	http.ListenAndServe(":5000", router)
}
