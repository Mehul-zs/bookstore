package serviceauthor

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
		err    error
	}{
		{desc: "Valid Author", input: entities.Author{10, "Mehul", "Abc", "20/02/1925", "Hey"},
			expOut: http.StatusCreated, err: nil},
		{"InValid Author", entities.Author{2, "Mehul", "Rawal", "20/02/1935", "Hie"},
			http.StatusBadRequest, errors.New(" Duplicate entry '1' for key 'Author.PRIMARY'")},
		{"Invalid Author", entities.Author{1, "", "Rawal", "20/02/1935", "Hie"},
			http.StatusBadRequest, errors.New("either of the field is empty")},
		{"Invalid Author", entities.Author{-3, "Mehul", "Rawal", "20/02/1935", "Hie"},
			http.StatusBadRequest, errors.New("either of the field is empty")},
	}

	a := New(mockdatastore{})

	for i, tc := range testcases {

		res, err := a.PostAuthor(tc.input)
		if res != tc.expOut && err != tc.err {
			t.Errorf("testcase:%d desc :%s actualError:%v actualResult:%v expectedError:%v expectedResult:%v", i, tc.desc, err, res, tc.err, tc.expOut)
		}

	}

}

func TestPutAuthor(t *testing.T) {
	testcases := []struct {
		desc   string
		id     int
		req    entities.Author
		expOut entities.Author
		err    error
	}{
		{desc: "Valid Author", id: 1, req: entities.Author{Id: 1, FirstName: "Mehul", LastName: "Rawal", Dob: "06/05/2001", PenName: "me"},
			expOut: entities.Author{Id: 1, FirstName: "Mehul", LastName: "Rawal", Dob: "06/05/2001", PenName: "me"}, err: nil},
		{"Invalid Author", 1, entities.Author{Id: 1, FirstName: "", LastName: "Rawal", Dob: "06/05/2001", PenName: "me"},
			entities.Author{}, errors.New("not valid constraints")},
		{"Inalid Author", 5, entities.Author{Id: 5, FirstName: "Mehul", LastName: "Rawal", Dob: "06/05/2001", PenName: "me"},
			entities.Author{}, errors.New("invalid author details")},
		{desc: "Invalid ID", id: -1, req: entities.Author{-1, "Mehul", "Rawal", "20/02/1935", "Hie"},
			expOut: entities.Author{}, err: errors.New("Negative ID params")},
	}

	for i, tc := range testcases {
		a := New(mockdatastore{})
		res, err := a.PutAuthor(tc.req, tc.id)
		if err != tc.err && tc.expOut != res {
			t.Errorf("testcase:%d desc :%s actualError:%v actualResult:%v expectedError:%v expectedResult:%v", i, tc.desc, err, res, tc.err, tc.expOut)
		}
	}
}

func TestDeleteAuthor(t *testing.T) {
	testcases := []struct {
		desc   string
		input  int
		expOut int64
		err    error
	}{
		{"Valid AuthorId", 1, http.StatusNoContent, nil},
		{"Invalid AuthorId", -2, http.StatusBadRequest, errors.New("negative id not accepted")},
		{"InValid AuthorId", 4, http.StatusBadRequest, errors.New("??")},

		//{"Invalid AuthorId", , http.StatusBadRequest, errors.New("not valid id")},
	}

	for i, tc := range testcases {
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

func (m mockdatastore) PostAuthor(author entities.Author) (int64, error) {
	if author.Id == 10 {
		return http.StatusCreated, nil
	}

	return http.StatusBadRequest, errors.New(" Duplicate entry '1' for key 'Author.PRIMARY'")
}

func (m mockdatastore) PutAuthor(author entities.Author, id int) (entities.Author, error) {
	if author.Id == 1 {
		return entities.Author{Id: 1, FirstName: "Mehul", LastName: "Rawal", Dob: "06/05/2001", PenName: "me"}, nil
	}

	return entities.Author{}, errors.New("invalid author details")
}

func (m mockdatastore) DeleteAuthor(id int) (int64, error) {
	if id == 1 {
		return http.StatusNoContent, nil
	}
	return http.StatusBadRequest, errors.New("??")
}
