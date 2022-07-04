package services

import "Bookstore/entities"

type Book interface {
	GetAll(books entities.Books) ([]entities.Books, error)
	GetByID(books entities.Books) (entities.Books, error)
	PostBook(books entities.Books) (int64, error)
	PutBook(books entities.Books) (int64, error)
	DeleteBook(int) (int64, error)
}

type Author interface {
	PostAuthor(author entities.Author) (int64, error)
	DeleteAuthor(int) (int64, error)
	PutAuthor(author entities.Author) (int64, error)
}
