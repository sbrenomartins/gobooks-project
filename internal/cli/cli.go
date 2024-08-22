package cli

import "github.com/sbrenomartins/gobooks/internal/service"

type BookCLI struct {
	service *service.BookService
}

func NewBookCLI(service *service.BookService) *BookCLI {
	return &BookCLI{service: service}
}
