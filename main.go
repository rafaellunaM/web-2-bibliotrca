package main

import (
	crud "library/crud"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var err error

func checErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books", crud.GetBooks).Methods("GET")
	router.HandleFunc("/books", crud.AddBook).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
