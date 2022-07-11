package handlerbook

import (
	"Bookstore/entities"
	"Bookstore/services"
	"bytes"
	"context"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
)

//func TestBookhandler_GetAll(t *testing.T) {
//
//	testcases := []struct {
//		desc   string
//		expOut []entities.Books
//	}{
//		{"Valid", []entities.Books{{
//			Id:     1,
//			Title:  "Jail",
//			Author: entities.Author{
//				//Id:        2,
//				//FirstName: "Charles",
//				//LastName:  "Lee",
//				//Dob:       "12/01/1942",
//				//PenName:   "CL",
//			},
//			Publication:   "Penguin",
//			PublishedDate: "12/02/1970",
//			AuthorID:      2,
//		},
//		},
//		},
//	}
//
//	for _, tc := range testcases {
//		req := httptest.NewRequest("GET", "/books", nil)
//		rw := httptest.NewRecorder()
//
//		a := New(mock{})
//
//		a.GetAll(rw, req)
//		data, err := io.ReadAll(rw.Body)
//		if err != nil {
//			t.Errorf("test case fail ,error in reading body")
//		}
//
//		var output []entities.Books
//
//		err = json.Unmarshal(data, &output)
//		if err != nil {
//			t.Errorf("test case fail ,error in unmarshaling data")
//		}
//
//		if !reflect.DeepEqual(output, tc.expOut) {
//			t.Errorf("test case failed")
//		}
//	}
//
//}

func TestBookhandler_GetByID(t *testing.T) {
	testcases := []struct {
		desc          string
		id            string
		expOut        entities.Book
		expStatusCode int
	}{

		{"invalid", "-1", entities.Book{}, http.StatusBadRequest},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tc := range testcases {
		mockServiceBook := services.NewMockBook(mockCtrl)
		b := New(mockServiceBook)
		req := httptest.NewRequest("GET", "/books/{id}="+tc.id, nil)
		req = mux.SetURLVars(req, map[string]string{"id": tc.id})

		rw := httptest.NewRecorder()
		id, _ := strconv.Atoi(v.id)
		mockServiceBook.EXPECT().GetBookByID(context.Background(), id).Return(v.expectedOutput, v.err).AnyTimes()
		var expout entities.Book
		b.GetBookByID(rw, req)
		body, err := io.ReadAll(rw.Body)
		if err != nil {
			log.Print(err)
		}
		err = json.Unmarshal(body, &expout)
		if err != nil {
			log.Print(err)
		}
		//if !reflect.DeepEqual(rw.Result().StatusCode, tc.expStatusCode) {
		//	t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i, rw.Result().StatusCode, tc.expStatusCode)
		//	//t.Errorf("Expected %v\tGot %v", v.expectedOutput, output)
		//}

	}

}

func TestBookhandler_PostBook(t *testing.T) {
	testcases := []struct {
		input     entities.Book
		expStatus int64
	}{
		{entities.Book{
			Id:    10,
			Title: "James",
			Author: entities.Author{
				Id:        2,
				FirstName: "Charles",
				LastName:  "Lee",
				Dob:       "12/01/1942",
				PenName:   "CL",
			},
			Publication:   "Penguin",
			PublishedDate: "12/02/1970",
			AuthorID:      2,
		}, http.StatusCreated,
		},
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for i, tc := range testcases {
		mockServiceBook := services.NewMockBook(mockCtrl)
		b := New(mockServiceBook)
		body, err := json.Marshal(tc.input)
		if err != nil {
			t.Errorf("Error converting in data format %v", tc.input)
		}
		req := httptest.NewRequest("POST", "/book", bytes.NewBuffer(body))
		rw := httptest.NewRecorder()

		mockServiceBook.EXPECT().PostBook(context.Background(), book)
		b.PostBook(rw, req)

		if !reflect.DeepEqual(rw.Result().StatusCode, v.expStatus) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i, rw.Result().StatusCode, v.expStatus)
		}
	}
}

func TestBookhandler_PutBook(t *testing.T) {
	testcases := []struct {
		desc          string
		input         entities.Book
		expout        entities.Book
		expstatuscode int64
		err           error
	}{
		{},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tc := range testcases {
		mockServiceBook := services.NewMockBook(mockCtrl)
		b := New(mockServiceBook)

		body, err := json.Marshal(tc.input)
		if err != nil {
			t.Errorf("Marshall error")
		}
		//req := httptest.NewRequest(http.MethodPut, "/book/{id}", bytes.NewBuffer(body))
		//book, err := tc.input.
	}

}

func TestBookhandler_DeleteBook(t *testing.T) {
	testcases := []struct {
		id        string
		expStatus int
		err       error
	}{
		{"abc", http.StatusBadRequest, "jss"},
		{"-1", http.StatusBadRequest, "jss"},
		{"1", http.StatusNoContent, nil},
		{"2", http.StatusNoContent, nil},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for i, tc := range testcases {
		mockServiceBook := services.NewMockBook(mockCtrl)
		b := New(mockServiceBook)

		req := httptest.NewRequest("DELETE", "/deleteBook/{id}"+tc.id, nil)
		req = mux.SetURLVars(req, map[string]string{"id": tc.id})
		ID, _ := strconv.Atoi(tc.id)
		rw := httptest.NewRecorder()

		mockServiceBook.EXPECT().DeleteBook(context.Background(), ID).Return(1, tc.err).AnyTimes()
		b.DeleteBook(rw, req)

		if !reflect.DeepEqual(rw.Result().StatusCode, v.expStatus) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i, rw.Result().StatusCode, tc.expStatus)
		}

	}

}
