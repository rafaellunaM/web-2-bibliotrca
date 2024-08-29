package handlers

import (
	"encoding/json"
	"library/pkg/models"
	"log"
	"net/http"
)

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {

	results, err := h.DB.Query("SELECT * FROM book;")
	if err != nil {
		log.Println("failed to execute query", err)
		w.WriteHeader(500)
		return
	}

	var books = make([]models.Book, 0)
	for results.Next() {
		var book models.Book
		err = results.Scan(&book.Bookid, &book.Title, &book.Author, &book.Quantity, &book.Category, &book.Price, &book.Availability)
		if err != nil {
			log.Println("failed to scan", err)
			w.WriteHeader(500)
			return
		}

		books = append(books, book)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
