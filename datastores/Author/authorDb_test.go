package datastoreAuthor

import (
	"Bookstore"
	"Bookstore/entities"
	"errors"
	"fmt"
	"testing"
)

func TestPostAuthor(t *testing.T) {

	testcases := []struct {
		desc string
		req  entities.Author
		err  error
	}{
		{"Valid Author", entities.Author{FirstName: "Mehul", LastName: "Rawal", Dob: "18/07/2000", PenName: "Me"}, errors.New("Sucess, status Ok!")},
		{"Author Already exists", entities.Author{1, "Rakshit", "Gupta", "12/02/1996", "kinu"}, errors.New("Author already exists")},
	}
	for _, tc := range testcases {

		DB, err := Bookstore.DbConn()
		if err != nil {
			fmt.Println("Db connection Error")
		}
		authorStore := New(DB)

		id, err := authorStore.PostAuthor(tc.req)

		if id == 0 && tc.err != err {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestDeleteAuthor(t *testing.T) {
	testcases := []struct {
		desc  string
		input int
		err   error
	}{
		{"Valid AuthorId", 4, nil},
		{"Invalid AuthorId", 5, errors.New("Invalid ID")},
	}

	for _, tc := range testcases {

		DB, err := Bookstore.DbConn()
		if err != nil {
			fmt.Println("Db connection Error")
		}
		authorStore := New(DB)

		id, err := authorStore.DeleteAuthor(tc.input)

		if id == 0 && tc.err != err {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}
