package main

import (
	"fmt"

	"github.com/sbrenomartins/gobooks/internal/service"
)

func main() {
	book := service.Book{
		ID:     1,
		Title:  "The Hobbit",
		Author: "J.R.R. Tolkien",
		Genre:  "Fantasy",
	}

	fmt.Println("the book is: ", book.GetFullBook())
}
