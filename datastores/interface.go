package datastore

import (
	"Bookstore/entities"
)

type BookStore interface {
	GetAll(books entities.Books) ([]entities.Books, error)
	GetByID(books entities.Books) (entities.Books, error)
	PostBook(books entities.Books) (int64, error)
	DeleteBook(books entities.Books) (int64, error)
	PutBook(books entities.Books) (int64, error)
}

type AuthorStore interface {
	PostAuthor(author entities.Author) (int64, error)
	DeleteAuthor(int) (int64, error)
	PutAuthor(author entities.Author) (int64, error)
}
