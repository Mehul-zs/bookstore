package datastoreBook

import (
	"Bookstore/entities"
	"database/sql"
)

type Bookstore struct {
	db *sql.DB
}

func (b Bookstore) DbConn() interface{} {

}

func New(db *sql.DB) Bookstore {
	return Bookstore{db: db}
}

func (b Bookstore) GetAll() ([]entities.Books, error) {
	return []entities.Books{}, nil
}

func (b Bookstore) GetByID(ID int) (entities.Books, error) {
	return entities.Books{}, nil
}

func (b Bookstore) PostBook(books entities.Books) (int64, error) {
	return 0, nil
}

func (b Bookstore) PutBook(books entities.Books) (int64, error) {
	return 0, nil
}

func (b Bookstore) DeleteBook(ID int) (int64, error) {
	return 0, nil
}
