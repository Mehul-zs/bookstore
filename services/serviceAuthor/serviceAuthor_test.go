package serviceAuthor

import (
	"Bookstore"
	"Bookstore/entities"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestPostAuthor(t *testing.T) (int, error) {
	testcases := []struct {
		desc   string
		input  entities.Author
		expOut int
	}{
		{"Valid Author", entities.Author{1, "Mehul", "Rawal", "20/02/1935", "Hie"}, http.StatusCreated},
	}
	for _, tc := range testcases {

		DB, err := Bookstore.DbConn()
		if err != nil {
			fmt.Println("database connection error")
		}
		authorStore := New(DB)
		authorService := New(authorStore)

		id, err := authorService.PostAuthor(tc.input)

		if err != nil {
			log.Print(err)
		}

		if id == 0 {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
	return 0, nil
}

func TestDeleteAuthor(t *testing.T) {
	testcases := []struct {
		desc  string
		input int
		//output
		expOut int
	}{
		{"Valid AuthorId", 2, http.StatusNoContent},
		{"Invalid AuthorId", 50, http.StatusBadRequest},
	}

	for _, tc := range testcases {
		DB, error := Bookstore.DbConn()
		authorStore := author.New(DB)
		authorService := New(authorStore)

		id, err := authorService.DeleteAuthor(tc.input)

		if err != nil {
			log.Print(err)
		}

		if id == 0 {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}
