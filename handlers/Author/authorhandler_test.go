package handlerAuthor

import (
	"Bookstore/entities"
	"bytes"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestAuthorHandler_Handler(t *testing.T) {
	testcases := []struct {
		method             string
		expectedStatusCode int
	}{
		{"POST", http.StatusOK},
		{"DELETE", http.StatusMethodNotAllowed},
	}

	for _, v := range testcases {
		req := httptest.NewRequest(v.method, "/author", nil)
		w := httptest.NewRecorder()

		a := New(mockDatastore{})
		a.Handler(w, req)

		if w.Code != v.expectedStatusCode {
			t.Errorf("Expected %v\tGot %v", v.expectedStatusCode, w.Code)
		}
	}
}

func TestPostAuthor(t *testing.T) {
	testcases := []struct {
		reqBody  []byte
		respBody []byte
	}{
		{[]byte(`{"FirstName":"Mehul","LastName":"", "Dob":"18/07/2000", "PenName":"Me"}`), []byte(`could not create animal`)},
		//{[]byte(`{"Name":"Maggie","Age":10}`), []byte(`{"ID":12,"Name":"Maggie","Age":10}`)},
		//{[]byte(`{"Name":"Maggie","Age":"10"}`), []byte(`invalid body`)},
	}
	for i, v := range testcases {
		req := httptest.NewRequest("GET", "/author", bytes.NewReader(v.reqBody))
		w := httptest.NewRecorder()

		a := New(mockDatastore{})

		a.PostAuthor(w, req)

		if !reflect.DeepEqual(w.Body, bytes.NewBuffer(v.respBody)) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Body.String(), string(v.respBody))
		}
	}
}

func DeleteAuthor(t *testing.T) {
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

		a.PostAuthor(w, req)




	}
}



func PutAuthor(t *testing.T){
	testcases:= []struct{
		input  []byte
		author []byte
		expStatus int
	}
}




type mockDatastore struct{}

func (m mockDatastore) PostAuthor(animal entities.Author) (entities.Author, error) {

	return entities.Author{}, nil

}
