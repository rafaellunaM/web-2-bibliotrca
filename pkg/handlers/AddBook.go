package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"library/pkg/models"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
		w.WriteHeader(500)
		return
	}
	var book models.Book
	json.Unmarshal(body, &book)

	queryStmt := `INSERT INTO book (bookid, title, author, quantity, category, price, availability) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING bookid;`
	err = h.DB.QueryRow(queryStmt, book.Bookid, book.Title, book.Author, book.Quantity, book.Category, book.Price, book.Availability).Scan(&book.Bookid)

	if err != nil {
		log.Println("failed to execute query", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")

}
