package models_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"yourapp/models"
	"net/http"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"yourapp/handlers"
	"github.com/jackc/pgx/v4"
)

/*
##########################
##											##
##  		unit tests  		##
##											##
##########################
*/

// Teste para validar a adição de um livro
func TestAddBook(t *testing.T) {
	book := models.Book{
		Bookid:       "12345",
		Title:        "Go Programming",
		Author:       "John Doe",
		Quantity:     10,
		Category:     "Programming",
		Price:        29.99,
		Availability: true,
	}

	// Simula a adição do livro (essa função deve ser implementada no seu código)
	err := addBookToDB(book)  // Função fictícia que você pode substituir pela implementação real.

	// Verifica se o erro é nil (sem erro)
	assert.Nil(t, err)

	// Verifica se o livro foi realmente adicionado (essa verificação depende da sua implementação)
	// Aqui você pode adicionar lógica para verificar no banco de dados se o livro foi realmente adicionado.
}

func TestUpdateBook(t *testing.T) {
	updatedBook := models.Book{
		Bookid:       "12345",
		Title:        "Advanced Go Programming",
		Author:       "John Doe",
		Quantity:     15,
		Category:     "Programming",
		Price:        39.99,
		Availability: true,
	}

	err := updateBookInDB(updatedBook)  // Função fictícia que deve ser implementada

	assert.Nil(t, err)

	// Verifique se o livro foi atualizado corretamente no banco de dados
}

func TestDeleteBook(t *testing.T) {
	bookID := "12345"
	err := deleteBookFromDB(bookID)  // Função fictícia que deve ser implementada

	assert.Nil(t, err)

	// Verifique se o livro foi deletado corretamente no banco de dados
}

func TestDBConnection(t *testing.T) {
	db, err := pgx.Open("postgres", "user=postgres dbname=mydb sslmode=disable")
	assert.Nil(t, err)
	defer db.Close()

	err = db.Ping()
	assert.Nil(t, err)
}

/*
##########################
##											##
##  integrations tests  ##
##											##
##########################
*/

func TestAddBookAPI(t *testing.T) {
	book := models.Book{
		Bookid:       "12345",
		Title:        "Go Programming",
		Author:       "John Doe",
		Quantity:     10,
		Category:     "Programming",
		Price:        29.99,
		Availability: true,
	}

	// Serializa o livro para JSON
	bookJSON, err := json.Marshal(book)
	if err != nil {
		t.Fatalf("Erro ao serializar livro: %v", err)
	}

	// Fazendo a requisição POST para a API
	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(bookJSON))
	if err != nil {
		t.Fatalf("Erro ao criar a requisição: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Testando a resposta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.AddBook)
	handler.ServeHTTP(rr, req)

	// Verifica se o status da resposta é 200 OK
	assert.Equal(t, http.StatusOK, rr.Code)

	// Aqui você pode verificar se o livro foi adicionado ao banco de dados.
}

func TestGetBookAPI(t *testing.T) {
	req, err := http.NewRequest("GET", "/books/12345", nil)
	if err != nil {
		t.Fatalf("Erro ao criar a requisição: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetBook)
	handler.ServeHTTP(rr, req)

	// Verifica se o status da resposta é 200 OK
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verifique o conteúdo da resposta aqui
	// Por exemplo, você pode decodificar o JSON e verificar os campos.
}

func TestUpdateBookAPI(t *testing.T) {
	book := models.Book{
		Bookid:       "12345",
		Title:        "Go Programming",
		Author:       "John Doe",
		Quantity:     10,
		Category:     "Programming",
		Price:        29.99,
		Availability: true,
	}
	addBookToDB(book)  // Adiciona o livro para garantir que ele existe no banco

	// Atualiza o livro
	updatedBook := models.Book{
		Bookid:       "12345",
		Title:        "Go Programming - Updated",
		Author:       "John Doe",
		Quantity:     12,
		Category:     "Programming",
		Price:        34.99,
		Availability: true,
	}

	bookJSON, err := json.Marshal(updatedBook)
	assert.Nil(t, err)

	req, err := http.NewRequest("PUT", "/books/12345", bytes.NewBuffer(bookJSON))
	assert.Nil(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.UpdateBook)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteBookAPI(t *testing.T) {
	book := models.Book{
		Bookid:       "12345",
		Title:        "Go Programming",
		Author:       "John Doe",
		Quantity:     10,
		Category:     "Programming",
		Price:        29.99,
		Availability: true,
	}
	addBookToDB(book)  // Adiciona o livro ao banco de dados

	req, err := http.NewRequest("DELETE", "/books/12345", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.DeleteBook)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	// Verifique se o livro foi realmente deletado
}

func TestGetAllBooksAPI(t *testing.T) {
	req, err := http.NewRequest("GET", "/books", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetAllBooks)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	// Aqui você pode verificar o conteúdo do JSON retornado
}

func TestAddBookInvalidPriceAPI(t *testing.T) {
	invalidBook := models.Book{
		Bookid:       "12345",
		Title:        "Go Programming",
		Author:       "John Doe",
		Quantity:     10,
		Category:     "Programming",
		Price:        -29.99, // Preço inválido
		Availability: true,
	}

	bookJSON, err := json.Marshal(invalidBook)
	assert.Nil(t, err)

	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(bookJSON))
	assert.Nil(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.AddBook)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}


func TestGetAllBooksDatabaseErrorAPI(t *testing.T) {
	// Simula a falha de conexão com o banco
	// Você pode fazer isso utilizando uma configuração de banco incorreta para esse teste específico

	req, err := http.NewRequest("GET", "/books", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetAllBooks)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
