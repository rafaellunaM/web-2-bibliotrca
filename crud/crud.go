package crud

import (
	"database/sql"
	"encoding/json"
	tables "library/dbconfig"
	"net/http"
)

var db *sql.DB
var err error

func GetBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM Book")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []tables.Book
	for rows.Next() {
		var book tables.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Quantity, &book.Price, &book.Availability); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var book tables.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO books (bookId, title, author, category, price, availability) VALUES (?, ?, ?, ?, ?, ?)",
		book.ID, book.Title, book.Author, book.Quantity, book.Price, book.Availability)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
