package handlerAuthor

import (
	"Bookstore/entities"
	"bytes"
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
		{[]byte(`{"FirstName":"Mehul","LastName":"", "Dob":"18/07/2000", "PenName":"Me"}`), []byte(`{"FirstName":"Mehul","LastName":"", "Dob":"18/07/2000", "PenName":"Me"}`)},
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
	for i, v := range testcases {
		req := httptest.NewRequest("DELETE", v.ID, nil)
		w := httptest.NewRecorder()

		a := New(mockDatastore{})

		a.DeleteAuthor(w, req)

		if !reflect.DeepEqual(w.Result().StatusCode, v.expStatus) {
			t.Errorf("Test case failed")
		}

	}
}

type mockDatastore struct{}

func (m mockDatastore) PostAuthor(a entities.Author) (entities.Author, error) {

	return entities.Author{}, nil

}

func (m mockDatastore) PutAuthor(a entities.Author) (int64, error) {
	return 0, nil
}

func (m mockDatastore) DeleteAuthor(id int) (int64, error) {
	return 0, nil
}
