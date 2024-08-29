package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"library/pkg/db"
	"library/pkg/handlers"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the library REST API!")
	fmt.Println("Library REST API")
}

func handleRequests(DB *sql.DB) {
	h := handlers.New(DB)
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	myRouter.HandleFunc("/books/{bookid}", h.GetBook).Methods(http.MethodGet)
	myRouter.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	myRouter.HandleFunc("/books/{bookid}", h.UpdatedBook).Methods(http.MethodPut)
	myRouter.HandleFunc("/books/{bookid}", h.DeleteBook).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
	fmt.Println("Listening in port 8080")
}

func main() {
	DB := db.Connect()
	handleRequests(DB)
	db.CloseConnection(DB)
}
