package serviceauthor

import (
	datastore "Bookstore/datastore"
	"Bookstore/entities"
	"context"
	"errors"
	"fmt"
	"net/http"
)

type ServiceAuthor struct {
	authorstore datastore.AuthorStore
}

func New(a datastore.AuthorStore) ServiceAuthor {
	return ServiceAuthor{a}
}

func (s ServiceAuthor) PostAuthor(ctx context.Context, a entities.Author) (int64, error) {

	if a.FirstName == "" || a.LastName == "" || a.Id <= 0 || a.Dob == "" || a.PenName == "" {
		return http.StatusBadRequest, errors.New("either of the field is empty")
	}

	chk, err := s.authorstore.CheckAuthor(ctx, a)
	if err != nil {
		return 0, err
	}

	if chk == 0 {
		_, err = s.authorstore.PostAuthor(ctx, a)
		if err != nil {
			return 0, err
		}
		return 1, nil
	}
	return 0, err
	//return 0, nil
}

func (s ServiceAuthor) PutAuthor(ctx context.Context, a entities.Author, id int) (entities.Author, error) {

	if a.FirstName == "" || a.LastName == "" || a.Id <= 0 || a.Dob == "" || a.PenName == "" {
		return entities.Author{}, errors.New("not valid constraints")
	}
	if id <= 0 {
		return entities.Author{}, errors.New("negative ID params")
	}
	_, err := s.authorstore.CheckAuthor(ctx, a)
	if err != nil {
		return entities.Author{}, err
	}

	author, err := s.authorstore.PutAuthor(ctx, a, id)
	if err != nil {
		return entities.Author{}, err
	}
	fmt.Println("Hello handler 4")

	return author, nil
	fmt.Println("Hello handler 10")

	//return entities.Author{}
	return entities.Author{}, err
}

// delete author -- completed
func (s ServiceAuthor) DeleteAuthor(ctx context.Context, id int) (int64, error) {
	if id < 0 {
		return 0, errors.New("negative id not accpeted")
	}

	row, err := s.authorstore.DeleteAuthor(ctx, id)
	if err != nil {
		return 0, errors.New("not valid id")
	}

	return row, nil
	//return 0, nil
}
