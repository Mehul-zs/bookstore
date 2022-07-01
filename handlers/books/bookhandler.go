package handlerBook

import (
	"Bookstore/services"
)

type Bookhandler struct {
	serviceBook services.Book
}

func New(s services.Book) Bookhandler {
	return Bookhandler{serviceBook: s}
}
