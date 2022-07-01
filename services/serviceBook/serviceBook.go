package serviceBook

import "Bookstore/datastores"

type serviceBook struct {
	bookstore datastore.Book
}

func New(b datastore.Book) serviceBook {
	return serviceBook{b}
}
