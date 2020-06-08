package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/idasilva/dtk-knowledge/app/news"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{content}",news.HandlerFakeFinder)
	fmt.Println("Serve on at port:8080")
	http.ListenAndServe(":8000",router)
}
