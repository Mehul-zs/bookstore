package services

import (
	"Bookstore/entities"
	"context"
)

//go:generate mockgen -source=interface.go -destination=mockInterface.go -package=services
type Book interface {
	GetAllBooks(ctx context.Context, a string, b string) ([]entities.Book, error)
	GetBookByID(ctx context.Context, id int) (entities.Book, error)
	PostBook(ctx context.Context, book entities.Book) (int64, error)
	PutBook(ctx context.Context, book entities.Book, id int) (entities.Book, error)
	DeleteBook(ctx context.Context, id int) (int64, error)
}

type Author interface {
	PostAuthor(ctx context.Context, author entities.Author) (int64, error)
	PutAuthor(ctx context.Context, author entities.Author, id int) (entities.Author, error)
	DeleteAuthor(ctx context.Context, id int) (int64, error)
}
