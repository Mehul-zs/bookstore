package serviceAuthor

import (
	datastore "Bookstore/datastores"
	"Bookstore/entities"
	"errors"
	"net/http"
)

type ServiceAuthor struct {
	authorstore datastore.AuthorStore
}

func New(a datastore.AuthorStore) ServiceAuthor {
	return ServiceAuthor{a}
}

//func checkDob(Dob string) bool {
//
//	return false
//}

func (s ServiceAuthor) PostAuthor(a entities.Author) (int64, error) {

	if a.FirstName == "" {
		return 0, errors.New("not valid constraints")
	}

	_, err := s.authorstore.PostAuthor(a)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusCreated, nil // check and think whether id is returned or http status request
	//return 0, nil
}

func (s ServiceAuthor) PutAuthor(a entities.Author) (entities.Author, error) {

	if a.FirstName == "" {
		return entities.Author{}, errors.New("not valid constraints")
	}

	author, err := s.authorstore.PutAuthor(a)
	if err != nil {
		return entities.Author{}, err
	}

	return author, nil
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
