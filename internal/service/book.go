package service

import "database/sql"

type Book struct {
	ID     int
	Title  string
	Author string
	Genre  string
}

func (b Book) GetFullBook() string {
	return b.Title + " by " + b.Author + " of genre " + b.Genre
}

type BookService struct {
	db *sql.DB
}

func NewBookService(db *sql.DB) *BookService {
	return &BookService{db: db}
}

func (service *BookService) CreateBook(book *Book) error {
	query := "INSERT INTO books (title, author, genre) VALUES (?,?,?)"

	result, err := service.db.Exec(query, book.Title, book.Author, book.Genre)

	if err != nil {
		return err
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		return err
	}

	book.ID = int(lastInsertId)
	return nil
}

func (service *BookService) GetBooks() ([]Book, error) {
	query := "SELECT id, title, author, genre FROM books"
	rows, err := service.db.Query(query)

	if err != nil {
		return nil, err
	}

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (service *BookService) GetBookByID(id int) (*Book, error) {
	query := "SELECT id, title, author, genre FROM books WHERE id = ?"
	row := service.db.QueryRow(query, id)

	var book Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (service *BookService) UpdateBook(book *Book) error {
	query := "UPDATE books SET title = ?, author = ?, genre = ? WHERE id = ?"
	_, err := service.db.Exec(query, book.Title, book.Author, book.Genre, book.ID)
	return err
}

func (service *BookService) DeleteBook(id int) error {
	query := "DELETE FROM books WHERE id = ?"
	_, err := service.db.Exec(query, id)
	return err
}
