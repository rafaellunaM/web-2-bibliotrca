package handlers

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

type Book struct {
	Bookid       string
	Title        string
	Author       string
	Quantity     int
	Category     string
	Price        float64
	Availability bool
}

type Handler struct {
	DB *sql.DB
}

func (h *Handler) AddBook(book Book) error {
	query := `INSERT INTO book (bookid, title, author, quantity, category, price, availability) 
    VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING bookid;`

	row := h.DB.QueryRow(query, book.Bookid, book.Title, book.Author, book.Quantity, book.Category, book.Price, book.Availability)

	if err := row.Scan(&book.Bookid); err != nil {
		fmt.Printf("Erro ao add o livro: %v\n", err)
		return err
	}

	return nil
}

func TestAddBook(t *testing.T) { // iteration 1
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery(`\s*INSERT INTO book \(\s*bookid, title, author, quantity, category, price, availability\s*\)\s*VALUES\s*\(\$1, \$2, \$3, \$4, \$5, \$6, \$7\)\s*RETURNING bookid;\s*`).
		WithArgs("1", "Test Book", "Test Author", 10, "Test Category", 19.99, true).
		WillReturnRows(sqlmock.NewRows([]string{"bookid"}).AddRow("1"))

	h := Handler{DB: db}

	book := Book{
		Bookid:       "1",
		Title:        "Test Book",
		Author:       "Test Author",
		Quantity:     10,
		Category:     "Test Category",
		Price:        19.99,
		Availability: true,
	}

	err = h.AddBook(book)
	assert.NoError(t, err)

	mock.ExpectationsWereMet()
}

func TestAddBookSQLFail(t *testing.T) { // iteration 2
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery(`\s*INSERT INTO book \(\s*bookid, title, author, quantity, category, price, availability\s*\)\s*VALUES\s*\(\$1, \$2, \$3, \$4, \$5, \$6, \$7\)\s*RETURNING bookid;\s*`).
		WithArgs("1", "Test Book", "Test Author", 10, "Test Category", 19.99, true).
		WillReturnError(fmt.Errorf("SQL error"))

	h := Handler{DB: db}

	book := Book{
		Bookid:       "1",
		Title:        "Test Book",
		Author:       "Test Author",
		Quantity:     10,
		Category:     "Test Category",
		Price:        19.99,
		Availability: true,
	}

	err = h.AddBook(book)
	assert.Error(t, err)

	mock.ExpectationsWereMet()
}

func TestAddBookInvalidData(t *testing.T) { // iteration 3
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery(`\s*INSERT INTO book \(\s*bookid, title, author, quantity, category, price, availability\s*\)\s*VALUES\s*\(\$1, \$2, \$3, \$4, \$5, \$6, \$7\)\s*RETURNING bookid;\s*`).
		WithArgs("", "", "", 0, "", 0.0, false).
		WillReturnError(fmt.Errorf("invalid input"))

	h := Handler{DB: db}

	book := Book{
		Bookid:       "",
		Title:        "",
		Author:       "",
		Quantity:     0,
		Category:     "",
		Price:        0.0,
		Availability: false,
	}

	err = h.AddBook(book)
	assert.Error(t, err)

	mock.ExpectationsWereMet()
}

func TestAddBookDuplicate(t *testing.T) { // iteration 4
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery(`\s*INSERT INTO book \(\s*bookid, title, author, quantity, category, price, availability\s*\)\s*VALUES\s*\(\$1, \$2, \$3, \$4, \$5, \$6, \$7\)\s*RETURNING bookid;\s*`).
		WithArgs("1", "Test Book", "Test Author", 10, "Test Category", 19.99, true).
		WillReturnError(fmt.Errorf("duplicate bookid"))

	h := Handler{DB: db}

	book := Book{
		Bookid:       "1",
		Title:        "Test Book",
		Author:       "Test Author",
		Quantity:     10,
		Category:     "Test Category",
		Price:        19.99,
		Availability: true,
	}

	err = h.AddBook(book)
	assert.Error(t, err)

	mock.ExpectationsWereMet()
}

func (h *Handler) UpdateBook(book Book) error {
	query := `UPDATE book SET title = $1, price = $2 WHERE bookid = $3;`

	_, err := h.DB.Exec(query, book.Title, book.Price, book.Bookid)
	if err != nil {
		fmt.Println("Error during update:", err)
		return fmt.Errorf("failed to update book: %v", err)
	}

	return nil
}

func TestUpdateBook(t *testing.T) { // iteration 5
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	mock.ExpectExec(`\s*UPDATE book SET title = \$1, price = \$2 WHERE bookid = \$3;`).
		WithArgs("Updated Title", 25.99, "1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	h := Handler{DB: db}

	book := Book{
		Bookid: "1",
		Title:  "Updated Title",
		Price:  25.99,
	}

	err = h.UpdateBook(book)
	assert.NoError(t, err)

	mock.ExpectationsWereMet()
}

func TestUpdateBookFail(t *testing.T) { // iteration 6
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	mock.ExpectExec(`\s*UPDATE book SET title = \$1, price = \$2 WHERE bookid = \$3;`).
		WithArgs("Updated Title", 25.99, "1").
		WillReturnError(fmt.Errorf("update failed"))

	h := Handler{DB: db}

	book := Book{
		Bookid: "1",
		Title:  "Updated Title",
		Price:  25.99,
	}

	err = h.UpdateBook(book)
	assert.Error(t, err)

	mock.ExpectationsWereMet()
}

func (h *Handler) DeleteBook(bookid string) error {
	query := `DELETE FROM book WHERE bookid = $1;`
	_, err := h.DB.Exec(query, bookid)

	if err != nil {
		fmt.Println("Error during delete:", err)
		return fmt.Errorf("failed to delete book: %v", err)
	}
	return nil
}

func TestDeleteBook(t *testing.T) { // iteration 7
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	mock.ExpectExec(`\s*DELETE FROM book WHERE bookid = \$1;`).
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	h := Handler{DB: db}

	err = h.DeleteBook("1")
	assert.NoError(t, err)

	mock.ExpectationsWereMet()
}

func TestDeleteBookFail(t *testing.T) { // iteration 8
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	mock.ExpectExec(`\s*DELETE FROM book WHERE bookid = \$1;`).
		WithArgs("1").
		WillReturnError(fmt.Errorf("delete failed"))

	h := Handler{DB: db}

	err = h.DeleteBook("1")
	assert.Error(t, err)

	mock.ExpectationsWereMet()
}

func (h *Handler) TestDBConnection() error {

	err := h.DB.QueryRow("SELECT 1").Scan()
	if err != nil {
		fmt.Println("Error during connection:", err)
		return fmt.Errorf("database connection failed: %v", err)
	}
	return nil
}

func TestDBConnection(t *testing.T) { // iteration 9
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery(`SELECT 1`).
		WillReturnError(fmt.Errorf("database connection failed"))

	h := Handler{DB: db}

	err = h.TestDBConnection()
	assert.Error(t, err)

	mock.ExpectationsWereMet()
}

func TestDBConnectionFail(t *testing.T) { // iteration 10
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT 1").WillReturnError(fmt.Errorf("database connection failed"))

	h := Handler{DB: db}

	err = h.TestDBConnection()
	assert.Error(t, err)
}
