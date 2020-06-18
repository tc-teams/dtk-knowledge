package main

import (
	"github.com/gorilla/mux"
	"github.com/idasilva/fakefinder-crawler/appae/news"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{content}", news.HandlerFakeFinder)
	http.ListenAndServe(":8080",router)
}
