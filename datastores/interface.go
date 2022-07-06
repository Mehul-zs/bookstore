package datastore

import (
	"Bookstore/entities"
)

type BookStore interface {
	GetAll(string, string) ([]entities.Books, error)
	GetByID(int) (entities.Books, error)
	PostBook(books entities.Books) (int64, error)
	PutBook(books entities.Books, id int) (entities.Books, error)
	DeleteBook(int) (int64, error)
}

type AuthorStore interface {
	PostAuthor(author entities.Author) (entities.Author, error)
	PutAuthor(author entities.Author) (entities.Author, error)
	DeleteAuthor(int) (int64, error)
}
