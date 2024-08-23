package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/sbrenomartins/gobooks/internal/cli"
	"github.com/sbrenomartins/gobooks/internal/service"
	"github.com/sbrenomartins/gobooks/internal/web"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	bookService := service.NewBookService(db)
	bookHandler := web.NewBookHandlers(bookService)

	if len(os.Args) > 1 && (os.Args[1] == "search" || os.Args[1] == "simulate") {
		bookCLI := cli.NewBookCLI(bookService)
		bookCLI.Run()
		return
	}

	router := http.NewServeMux()

	router.HandleFunc("GET /books", bookHandler.GetAllBooks)
	router.HandleFunc("POST /books", bookHandler.CreateBook)
	router.HandleFunc("GET /books/{id}", bookHandler.GetBookByID)
	router.HandleFunc("PUT /books/{id}", bookHandler.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", bookHandler.DeleteBook)

	http.ListenAndServe(":8080", router)
}
