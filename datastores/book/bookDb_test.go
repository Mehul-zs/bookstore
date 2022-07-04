package datastoreBook

import (
	"Bookstore/entities"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestBookstore_GetAll(t *testing.T) {
	testcases := []struct {
		desc   string
		output []entities.Books
	}{
		{"Validated", []entities.Books{{
			Id:    1,
			Title: "James",
			Author: entities.Author{
				Id:        1,
				FirstName: "Mehul",
				LastName:  "Gupta",
				Dob:       "12/02/1970",
				PenName:   "Me",
			},
			Publication:   "Penguin",
			PublishedDate: "12/07/1999",
		}},
		},
	}

	for _, tc := range testcases {
		DB, err := Bookstore.DbConn()
		if err != nil {
			fmt.Println("Db connection Error")
		}
		bookStore := New(DB)

		chk, err := bookStore.GetAll()

		if err != nil {
			t.Errorf("Test case failed %v", tc.desc)
		}

		if !reflect.DeepEqual(chk, tc.output) {
			t.Errorf("Test case failed; %s", err)
		}
	}
}

func TestBookstore_GetByID(t *testing.T) {
	testcases := []struct {
		desc   string
		input  int
		expOut entities.Books
		err    error
	}{
		{"Valid book Id", 1, entities.Books{1, "james", entities.Author{FirstName: "Mehul", LastName: "Rawal", Dob: "18/07/2000", PenName: "Me"}, "Penguin", "12/01/2020"}, errors.New("Sucess, status Ok!")},
		{"Invalid book Id", -2, entities.Books{}, errors.New("invalid ID")},
	}
	for _, tc := range testcases {

		DB := Bookstore.DbConn()
		if err != nil {
			fmt.Println("Db connection Error")
		}
		bookStore := New(DB)

		chk, err := bookStore.GetByID(tc.input)

		if chk != tc.expOut {
			t.Errorf("Test case failed; %s", err)
		}
	}

}

func TestBookstore_PostBook(t *testing.T) {
	testcases := []struct {
		desc string
		req  entities.Books
		err  error
	}{
		{"Valid book", entities.Books{1, "james ", entities.Author{FirstName: "Mehul", LastName: "Rawal", Dob: "18/07/2000", PenName: "Me"}, "Penguin", "12/01/2020"}, errors.New("Sucess, status Ok!")},
		{"Invalid published date", entities.Books{1, "Rakshit", entities.Author{FirstName: "Mehul", LastName: "Rawal", Dob: "18/07/2000", PenName: "Me"}, "Penguin", "12/02/2222"}, errors.New("Author already exists")},
	}
	for _, tc := range testcases {

		DB, err := Bookstore.DbConn()
		if err != nil {
			fmt.Println("Db connection Error")
		}
		bookStore := New(DB)

		id, err := bookStore.PostBook(tc.req)

		if id == http.StatusBadRequest {
			t.Errorf("Bad request , test case failed")

		}
	}

}

func TestBookstore_DeleteBook(t *testing.T) {

	testcases := []struct {
		desc  string
		input int
		err   error
	}{
		{"Valid BookId", 4, nil},
		{"Invalid BookId", 5, errors.New("Invalid ID")},
	}

	for _, tc := range testcases {

		DB, err := Bookstore.DbConn()
		if err != nil {
			fmt.Println("Db connection Error")
		}
		authorStore := New(DB)

		id, err := authorStore.DeleteBook(tc.input)

		if id == 0 && tc.err != err {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}

}
