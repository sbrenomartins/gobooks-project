package web

import (
	"encoding/json"
	"net/http"

	"github.com/sbrenomartins/gobooks/internal/service"
)

type BookHandlers struct {
	bookService *service.BookService
}

func (handler *BookHandlers) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := handler.bookService.GetBooks()

	if err != nil {
		http.Error(w, "failed to get all books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
