package serviceBook

import (
	"Bookstore/entities"
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestServiceBook_GetAll(t *testing.T) {
	testcases := []struct {
		input  entities.Books
		output []entities.Books
	}{}

	for _, tc := range testcases {

	}

}

func TestServiceBook_GetByID(t *testing.T) {
	testcases := []struct {
		input  int
		output entities.Books
	}{
		{},
	}

	for _, tc := range testcases {

		a := New(mockDatastore{})

		ans, err := a.GetByID(tc.input)
		if err != nil {
			t.Errorf("Test case failed %s", tc.input)
		}

		if ans != tc.output {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", tc.output, ans)
		}

	}
}

func TestServiceBook_PostBook(t *testing.T) {

	testcases := []struct {
		desc string
		req  entities.Books
		err  error
	}{
		{"Valid book", entities.Books{1, "james ", entities.Author{FirstName: "Mehul", LastName: "Rawal", Dob: "18/07/2000", PenName: "Me"}, "Penguin", "12/01/2020"}, errors.New("Sucess, status Ok!")},
		{"Invalid published date", entities.Books{1, "Rakshit", entities.Author{FirstName: "Mehul", LastName: "Rawal", Dob: "18/07/2000", PenName: "Me"}, "Penguin", "12/02/2222"}, errors.New("Author already exists")},
	}
	for _, tc := range testcases {

		DB, err := main.DbConn()
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

func TestServiceBook_DeleteBook(t *testing.T) {
	testcases := []struct {
		desc  string
		input int
		err   error
	}{
		{"Valid BookId", 4, nil},
		{"Invalid BookId", 5, errors.New("Invalid ID")},
	}

	for _, tc := range testcases {

		DB, err := main.DbConn()
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

type mockDatastore struct{}

func (m mockDatastore) getAll(books entities.Books) (entities.Books, error) {

}

func (m mockDatastore) getById(id int) (entities.Books, error) {

}

func (m mockDatastore) postBook(books entities.Books) (entities.Books, error) {

}

func (m mockDatastore) putBook(books entities.Books) (entities.Books, error) {

}

func (m mockDatastore) deleteBook(id int) (entities.Books, error) {

}
