package serviceAuthor

import (
	"Bookstore/entities"
	"errors"
	"log"
	"net/http"
	"testing"
)

func TestPostAuthor(t *testing.T) {
	testcases := []struct {
		desc   string
		input  entities.Author
		expOut int64
	}{
		{"Valid Author", entities.Author{1, "Mehul", "Rawal", "20/02/1935", "Hie"}, http.StatusCreated},
	}

	a := New(mockdatastore{})

	for i, tc := range testcases {

		res, err := a.PostAuthor(tc.input)
		if res != tc.expOut {
			t.Errorf("testcase:%d desc :%s actualError:%v actualResult:%v expectedError:%v expectedResult:%v", i, tc.desc, res, err, tc.expOut)
		}

	}

}

func TestDeleteAuthor(t *testing.T) {
	testcases := []struct {
		desc  string
		input int
		//output
		expOut int64
		err    error
	}{
		{"Valid AuthorId", 2, http.StatusNoContent, nil},
		{"Invalid AuthorId", 50, http.StatusBadRequest, errors.New("Author not present")},
	}

	for i, tc := range testcases {

		//id, err := authorService.DeleteAuthor(tc.input)
		a := New(mockdatastore{})
		res, err := a.DeleteAuthor(tc.input)
		if err != nil {
			log.Print(err)
		}

		if res != tc.expOut && err != tc.err {
			t.Errorf("testcase:%d desc:%v actualResult:%v actualError:%v expectedResponse:%v expectedError:%v", i, tc.desc, res, err, tc.expOut, tc.err)
		}

	}
}

type mockdatastore struct{}

func (m mockdatastore) PostAuthor(author entities.Author) (entities.Author, error) {
	if author.Id == 1 {
		return entities.Author{1, "Mehul", "Rawal", "20/02/1935", "Hie"}, nil
	}

	return entities.Author{}, errors.New(" Duplicate entry '1' for key 'Author.PRIMARY key'")
}

func (m mockdatastore) PutAuthor(author entities.Author) (entities.Author, error) {
	return entities.Author{}, nil
}

func (m mockdatastore) DeleteAuthor(id int) (int64, error) {
	if id == 1 {
		return 1, nil
	}
	return 0, nil
}
