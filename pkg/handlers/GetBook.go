package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"library/pkg/models"

	"github.com/gorilla/mux"
)

func (h handler) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["bookid"]

	queryStmt := `SELECT * FROM book WHERE bookid = $1 ;`
	results, err := h.DB.Query(queryStmt, id)
	if err != nil {
		log.Println("failed to execute query", err)
		w.WriteHeader(500)
		return
	}

	var book models.Book
	for results.Next() {
		err = results.Scan(&book.Bookid, &book.Title, &book.Author, &book.Quantity, &book.Category, &book.Price, &book.Availability)
		if err != nil {
			log.Println("failed to scan", err)
			w.WriteHeader(500)
			return
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
