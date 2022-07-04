package datastoreAuthor

import (
	"Bookstore/entities"
	"database/sql"
)

type Authorstore struct {
	db *sql.DB
}

func New(db *sql.DB) Authorstore {
	return Authorstore{db: db}
}

func (a Authorstore) PostAuthor(author entities.Author) (int, error) {
	return 0, nil
}

func (a Authorstore) PutAuthor(author entities.Author) (int, error) {
	return 0, nil
}

func (a Authorstore) DeleteAuthor(id int) (int, error) {

	return 0, nil

}
