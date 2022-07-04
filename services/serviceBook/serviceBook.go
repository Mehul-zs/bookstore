package serviceBook

import (
	datastore "Bookstore/datastores"
	datastoreBook "Bookstore/datastores/book"
	"Bookstore/entities"
)

type serviceBook struct {
	bookstore datastore.BookStore
}

func New(b datastoreBook.Bookstore) serviceBook {
	return serviceBook{b}
}

func (bs serviceBook) GetAll(books entities.Books) ([]entities.Books, error) {
	return nil, nil
}

func (bs serviceBook) GetByID(id int) (entities.Books, error) {
	return entities.Books{}, nil
}

func (bs serviceBook) PostBook(books entities.Books) (int64, error) {
	return 0, nil
}

func (bs serviceBook) PutBook(books entities.Books) (int64, error) {

	return 0, nil
}

func (bs serviceBook) DeleteBook(int) (int64, error) {
	return 0, nil
}
