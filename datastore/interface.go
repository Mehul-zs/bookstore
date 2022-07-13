package datastore

import (
	"Bookstore/entities"
	"context"
)

//go:generate mockgen -destination=mockInterface.go -package=datastore -source=interface.go

type BookStore interface {
	GetAllBooks(ctx context.Context, a string, b string) ([]entities.Book, error)
	GetAllBooksByTitle(ctx context.Context, title string) ([]entities.Book, error)
	GetBookByID(ctx context.Context, id int) (entities.Book, error)
	PostBook(ctx context.Context, books *entities.Book) (int64, error)
	PutBook(ctx context.Context, books entities.Book, id int) (entities.Book, error)
	DeleteBook(ctx context.Context, id int) (int64, error)
	CheckBook(ctx context.Context, id int) (bool, error)
}

type AuthorStore interface {
	GetAllAuthor(ctx context.Context) ([]entities.Author, error)
	GetAuthorByID(ctx context.Context, id int) (entities.Author, error)
	CheckAuthor(ctx context.Context, author entities.Author) (int, error)
	CheckAuthorByID(ctx context.Context, id int) (bool, error)
	PostAuthor(ctx context.Context, author entities.Author) (int64, error)
	PutAuthor(ctx context.Context, author entities.Author, id int) (entities.Author, error)
	DeleteAuthor(ctx context.Context, id int) (int64, error)
}
