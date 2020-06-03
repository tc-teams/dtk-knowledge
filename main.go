package main

import (
	"github.com/gorilla/mux"
	"github.com/idasilva/dtk-knowledge/app/news"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/",news.HandlerFakeFinder)
	http.ListenAndServe(":8080",router)
}
