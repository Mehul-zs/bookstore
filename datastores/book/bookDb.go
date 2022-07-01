package datastoreBook

import (
	"Bookstore/entities"
	"database/sql"
)

type Bookstore struct {
	db *sql.DB
}

func (b Bookstore) GetAll() ([]entities.Books, error) {
	return nil
}

func (b Bookstore) GetByID(ID int) (entities.Books, error) {
	return
}

func (b Bookstore) PostBook() {

}
