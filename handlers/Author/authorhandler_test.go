package handlerauthor

import (
	"Bookstore/services"
	"bytes"
	"errors"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

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

func TestPutAuthor(t *testing.T) {

}

func TestDeleteAuthor(t *testing.T) {
	testcases := []struct {
		desc      string
		ID        string
		expStatus int
		err       error
	}{
		//{"Valid", "50", http.StatusBadRequest, },
		{"Invalid input", "-1", http.StatusBadRequest, errors.New("invalid id")},
		{"Valid inout", "1", http.StatusNoContent, nil},
	}

	mockControl := gomock.NewController(t)
	defer mockControl.Finish()

	for i, tc := range testcases {
		mockServiceAuthor := services.NewMock

		req := httptest.NewRequest("DELETE", "/author/{id}"+tc.ID, nil)
		rw := httptest.NewRecorder()

		a.DeleteAuthor(rw, req)

		if !reflect.DeepEqual(rw.Result().StatusCode, tc.expStatus) {
			t.Errorf("testcase:%d desc:%v actualResult:%v expectedResponse:%v expectedError:%v", i, tc.desc, rw.Result().StatusCode, tc.expStatus, tc.err)
		}

	}
}
