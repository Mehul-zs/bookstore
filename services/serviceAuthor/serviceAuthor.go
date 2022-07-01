package serviceAuthor

import (
	datastore "Bookstore/datastores"
	"Bookstore/entities"
)

type ServiceAuthor struct {
	authorstore datastore.Author
}

func New(author datastore.Author) ServiceAuthor {
	return ServiceAuthor{authorstore: author}
}

func (s ServiceAuthor) PostAuthor() int {
	return 0
}

func (s ServiceAuthor) PutAuthor() entities.Author {
	return entities.Author{}
}

func (s ServiceAuthor) DeleteAuthor() int {
	return 0
}
