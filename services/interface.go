package services

import "Bookstore/entities"

type Book interface {
	GetAll(string, string) ([]entities.Books, error)
	GetByID(int) (entities.Books, error)
	PostBook(book entities.Books) (int64, error)
	PutBook(book entities.Books, id int) (entities.Books, error)
	DeleteBook(int) (int64, error)
}

type Author interface {
	PostAuthor(author entities.Author) (int64, error)
	PutAuthor(author entities.Author, id int) (entities.Author, error)
	DeleteAuthor(int) (int64, error)
}
