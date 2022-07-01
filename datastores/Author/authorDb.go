package datastoreAuthor

import "database/sql"

type Authorstore struct {
	db *sql.DB
}

func New(db *sql.DB) Authorstore {
	return Authorstore{db: db}
}

func (a Authorstore) PostAuthor() {

}

func (a Authorstore) PutAuthor() {

}

func (a Authorstore) DeleteAuthor() {

}
