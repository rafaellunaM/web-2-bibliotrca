package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"library/pkg/models"

	"github.com/gorilla/mux"
)

func (h handler) UpdatedBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["bookid"]

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)

	queryStmt := `UPDATE book SET title = $2, author = $3, quantity = $4, category = $5, price = $6, availability = $7 WHERE bookid = $1 RETURNING bookid;`
	err = h.DB.QueryRow(queryStmt, &updatedBook.Bookid, &updatedBook.Title, &updatedBook.Author, &updatedBook.Quantity, &updatedBook.Category, &updatedBook.Price, &updatedBook.Availability).Scan(&id)
	if err != nil {
		log.Println("failed to execute query", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")

}
