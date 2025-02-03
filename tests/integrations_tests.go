package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"yourapp/handlers"
	"yourapp/models"
	"github.com/stretchr/testify/assert"
)

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
