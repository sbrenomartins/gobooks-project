package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sbrenomartins/gobooks/internal/service"
)

type BookHandlers struct {
	bookService *service.BookService
}

func NewBookHandlers(bookService *service.BookService) *BookHandlers {
	return &BookHandlers{bookService: bookService}
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

func (handler *BookHandlers) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book service.Book

	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	err = handler.bookService.CreateBook(&book)

	if err != nil {
		http.Error(w, "failed to create book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (handler *BookHandlers) GetBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := handler.bookService.GetBookByID(id)

	if err != nil {
		http.Error(w, "failed to get book", http.StatusInternalServerError)
		return
	}

	if book == nil {
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (handler *BookHandlers) UpdateBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "invalid book ID", http.StatusBadRequest)
		return
	}

	var book service.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	book.ID = id

	if err := handler.bookService.UpdateBook(&book); err != nil {
		http.Error(w, "failed to update book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (handler *BookHandlers) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "invalid book ID", http.StatusBadRequest)
		return
	}

	if err := handler.bookService.DeleteBook(id); err != nil {
		http.Error(w, "failed to delete book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
