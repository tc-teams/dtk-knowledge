package main

import (
	"github.com/gorilla/mux"
	"github.com/idasilva/fakefinder-crawler/app/news"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{content}", news.HandlerFakeFinder)
	http.ListenAndServe(":8080",router)
}
