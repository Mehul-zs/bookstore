package handlerBook

import (
	"Bookstore/entities"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestBookhandler_GetAll(t *testing.T) {

	testcases := []struct {
		id     string
		expOut entities.Books
	}{
		{"1", entities.Books{
			Id:    1,
			Title: "Mehul",
			Author: entities.Author{
				Id:        0,
				FirstName: "Charles",
				LastName:  "Lee",
				Dob:       "12/01/1942",
				PenName:   "CL",
			},
			Publication:   "Penguin",
			PublishedDate: "12/02/1970",
		}},
	}

	for i, v := range testcases {
		req := httptest.NewRequest("GET", "/animal?id="+v.id, nil)
		w := httptest.NewRecorder()

		a := New(mockDatastore{})

		a.GetAll(w, req)

		if !reflect.DeepEqual(w.Body, v.expOut) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Body.String(), v.expOut)
		}
	}

}

func TestBookhandler_GetByID(t *testing.T) {
	testcases := []struct {
		id     string
		expOut entities.Books
	}{
		{"1", entities.Books{
			Id:    1,
			Title: "Mehul",
			Author: entities.Author{
				Id:        0,
				FirstName: "Charles",
				LastName:  "Lee",
				Dob:       "12/01/1942",
				PenName:   "CL",
			},
			Publication:   "Penguin",
			PublishedDate: "12/02/1970",
		}},
	}

	for i, v := range testcases {
		req := httptest.NewRequest("GET", "/book?id="+v.id, nil)
		w := httptest.NewRecorder()

		a := New(mockDatastore{})

		a.GetAll(w, req)

		if !reflect.DeepEqual(w.Body, v.expOut) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Body.String(), v.expOut)
		}
	}
}

func TestBookhandler_PostBook(t *testing.T) {

	testcases := []struct {
		input  entities.Books
		expOut entities.Books
	}{
		{entities.Books{
			Id:    1,
			Title: "Mehul",
			Author: entities.Author{
				Id:        0,
				FirstName: "Charles",
				LastName:  "Lee",
				Dob:       "12/01/1942",
				PenName:   "CL",
			},
			Publication:   "Penguin",
			PublishedDate: "12/02/1970",
		}, entities.Books{
			Id:    1,
			Title: "Mehul",
			Author: entities.Author{
				Id:        0,
				FirstName: "Charles",
				LastName:  "Lee",
				Dob:       "12/01/1942",
				PenName:   "CL",
			},
			Publication:   "Penguin",
			PublishedDate: "12/02/1970",
		},
		},
	}
	for i, v := range testcases {
		req := httptest.NewRequest("POST", "/book", nil)
		w := httptest.NewRecorder()

		a := New(mockDatastore{})

		a.PostBook(w, req)

		if !reflect.DeepEqual(w.Body, v.expOut) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Body.String(), v.expOut)
		}
	}

}

func TestBookhandler_DeleteBook(t *testing.T) {
	testcases := []struct {
		ID        string
		expStatus int
	}{
		{"50", http.StatusBadRequest},
		{"-1", http.StatusBadRequest},
		{"1", http.StatusNoContent},
	}
	for _, v := range testcases {
		req := httptest.NewRequest("DELETE", v.ID, nil)
		w := httptest.NewRecorder()

		a := New(mockDatastore{})

		a.PostBook(w, req)

		if !reflect.DeepEqual(w.Result().StatusCode, v.expStatus) {
			t.Errorf("Test case failed")
		}

	}

}

type mockDatastore struct{}

func (m mockDatastore) GetAll() {

}

func (m mockDatastore) GetByID() {

}

func (m mockDatastore) PostBook() {

}

func (m mockDatastore) PutBook() {

}

func (m mockDatastore) DeleteBook() {

}
