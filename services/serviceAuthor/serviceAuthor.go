package serviceAuthor

import (
	datastore "Bookstore/datastores"
	datastoreAuthor "Bookstore/datastores/Author"
	"Bookstore/entities"
	"errors"
)

type ServiceAuthor struct {
	authorstore datastore.AuthorStore
}

func New(author datastoreAuthor.Authorstore) ServiceAuthor {
	return ServiceAuthor{authorstore: author}
}

func checkDob(Dob string) bool {
	return false
}

func (s ServiceAuthor) PostAuthor(a entities.Author) (int64, error) {

	if a.FirstName == "" || !checkDob(a.Dob) {
		return 0, errors.New("not valid constraints")
	}

	id, err := s.authorstore.PostAuthor(a)
	if err != nil {
		return 0, err
	}

	return id, nil
	//return 0, nil
}

func (s ServiceAuthor) PutAuthor(a entities.Author) (int64, error) {
	if a.FirstName == "" || !checkDob(a.Dob) {
		return 0, errors.New("not valid constraints")
	}

	id, err := s.authorstore.PostAuthor(a)
	if err != nil {
		return 0, err
	}

	return id, nil
	//return entities.Author{}
}

func (s ServiceAuthor) DeleteAuthor(id int) (int64, error) {
	if id < 0 {
		return 0, errors.New("not valid id")
	}

	cnt, err := s.authorstore.DeleteAuthor(id)
	if err != nil {
		return 0, err
	}

	return cnt, nil
	//return 0, nil
}
