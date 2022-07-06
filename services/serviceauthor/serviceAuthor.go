package serviceauthor

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

func (s ServiceAuthor) PostAuthor(a entities.Author) (int64, error) {

	if a.FirstName == "" || a.LastName == "" || a.Id <= 0 || a.Dob == "" || a.PenName == "" {
		return http.StatusBadRequest, errors.New("either of the field is empty")
	}

	_, err := s.authorstore.PostAuthor(a)
	if err != nil {
		return http.StatusBadRequest, errors.New("")
	}

	return http.StatusCreated, nil // check and think whether id is returned or http status request
	//return 0, nil
}

func (s ServiceAuthor) PutAuthor(a entities.Author, id int) (entities.Author, error) {

	if a.FirstName == "" || a.LastName == "" || a.Id <= 0 || a.Dob == "" || a.PenName == "" {
		return entities.Author{}, errors.New("not valid constraints")
	}
	if id <= 0 {
		return entities.Author{}, errors.New("negative ID params")
	}

	author, err := s.authorstore.PutAuthor(a, id)
	if err != nil {
		return entities.Author{}, err
	}

	return author, nil
	//return entities.Author{}
}

func (s ServiceAuthor) DeleteAuthor(id int) (int64, error) {
	if id < 0 {
		return http.StatusBadRequest, errors.New("negative id not accpeted")
	}

	_, err := s.authorstore.DeleteAuthor(id)
	if err != nil {
		return http.StatusBadRequest, errors.New("not valid id")
	}

	return http.StatusNoContent, nil
	//return 0, nil
}
