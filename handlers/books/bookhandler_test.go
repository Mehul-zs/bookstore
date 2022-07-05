package handlerbook

import (
	"Bookstore/entities"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestBookhandler_GetAll(t *testing.T) {

	testcases := []struct {
		desc   string
		expOut []entities.Books
	}{
		{"Valid", []entities.Books{{
			Id:     1,
			Title:  "Jail",
			Author: entities.Author{
				//Id:        2,
				//FirstName: "Charles",
				//LastName:  "Lee",
				//Dob:       "12/01/1942",
				//PenName:   "CL",
			},
			Publication:   "Penguin",
			PublishedDate: "12/02/1970",
			AuthorID:      2,
		},
		},
		},
	}

	for _, tc := range testcases {
		req := httptest.NewRequest("GET", "/books", nil)
		rw := httptest.NewRecorder()

		a := New(mock{})

		a.GetAll(rw, req)
		data, err := io.ReadAll(rw.Body)
		if err != nil {
			t.Errorf("test case fail ,error in reading body")
		}

		var output []entities.Books

		err = json.Unmarshal(data, &output)
		if err != nil {
			t.Errorf("test case fail ,error in unmarshaling data")
		}

		if !reflect.DeepEqual(output, tc.expOut) {
			t.Errorf("test case failed")
		}
	}

}

func TestBookhandler_GetByID(t *testing.T) {
	testcases := []struct {
		desc          string
		id            string
		expOut        entities.Books
		expStatusCode int
	}{
		{"Valid", "1", entities.Books{
			Id:    1,
			Title: "Mehul",
			Author: entities.Author{
				Id:        1,
				FirstName: "Charles",
				LastName:  "Lee",
				Dob:       "12/01/1942",
				PenName:   "CL",
			},
			Publication:   "Penguin",
			PublishedDate: "12/02/1970",
			AuthorID:      1,
		}, http.StatusOK},
		{"invalid", "-1", entities.Books{}, http.StatusBadRequest},
	}

	for _, tc := range testcases {
		req := httptest.NewRequest("GET", "/books/{id}="+tc.id, nil)
		req = mux.SetURLVars(req, map[string]string{"id": tc.id})

		rw := httptest.NewRecorder()
		a := New(mock{})

		a.GetByID(rw, req)
		if rw.Result().StatusCode != http.StatusOK {
			if rw.Result().StatusCode != tc.expStatusCode {
				t.Errorf("test case fail")
			}
		} else {
			data, err := io.ReadAll(rw.Body)
			if err != nil {
				t.Errorf("test case fail ,error in reading body")
			}

			var output entities.Books

			err = json.Unmarshal(data, &output)
			if err != nil {
				t.Errorf("test case fail ,error in unmarshaling data")
			}

			if !reflect.DeepEqual(output, tc.expOut) {
				t.Errorf("test case failed")
				//t.Errorf("Expected %v\tGot %v", v.expectedOutput, output)
			}

		}
	}
}

//func TestBookhandler_PostBook(t *testing.T) {
//
//	testcases := []struct {
//		input  entities.Books
//		expOut entities.Books
//	}{
//		{entities.Books{
//			Id:    1,
//			Title: "Mehul",
//			Author: entities.Author{
//				Id:        0,
//				FirstName: "Charles",
//				LastName:  "Lee",
//				Dob:       "12/01/1942",
//				PenName:   "CL",
//			},
//			Publication:   "Penguin",
//			PublishedDate: "12/02/1970",
//		}, entities.Books{
//			Id:    1,
//			Title: "Mehul",
//			Author: entities.Author{
//				Id:        0,
//				FirstName: "Charles",
//				LastName:  "Lee",
//				Dob:       "12/01/1942",
//				PenName:   "CL",
//			},
//			Publication:   "Penguin",
//			PublishedDate: "12/02/1970",
//		},
//		},
//	}
//	for i, v := range testcases {
//		req := httptest.NewRequest("POST", "/book", nil)
//		w := httptest.NewRecorder()
//
//		//a := New(mockDatastore{})
//
//		a.PostBook(w, req)
//
//		if !reflect.DeepEqual(w.Body, v.expOut) {
//			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Body.String(), v.expOut)
//		}
//	}
//
//}

//func TestBookhandler_DeleteBook(t *testing.T) {
//	testcases := []struct {
//		ID        string
//		expStatus int
//	}{
//		{"50", http.StatusBadRequest},
//		{"-1", http.StatusBadRequest},
//		{"1", http.StatusNoContent},
//		{"2", http.StatusNoContent},
//	}
//	for _, v := range testcases {
//		req := httptest.NewRequest("DELETE", "/deleteBook/{id}"+v.ID, nil)
//		req = mux.SetURLVars(req, map[string]string{"id": v.ID})
//
//		rw := httptest.NewRecorder()
//		a := New(mock{})
//		a.DeleteBook(rw, req)
//
//		if !reflect.DeepEqual(rw.Result().StatusCode, v.expStatus) {
//			t.Errorf("Test case failed")
//		}
//
//	}
//
//}

type mock struct{}

//func (m mock) GetAll(s string, a string) ([]entities.Books, error) {
//	return nil, nil
//}
func (m mock) GetAll(title string, getAuthor string) ([]entities.Books, error) {
	return []entities.Books{
		{
			Id:     1,
			Title:  "Jail",
			Author: entities.Author{
				//Id:        2,
				//FirstName: "Charles",
				//LastName:  "Lee",
				//Dob:       "12/01/1942",
				//PenName:   "CL",
			},
			Publication:   "Penguin",
			PublishedDate: "12/02/1970",
			AuthorID:      2,
		}}, nil
}

func (m mock) GetByID(id int) (entities.Books, error) {
	if id <= 0 {
		return entities.Books{}, errors.New("invalid")
	}
	return entities.Books{
		Id:    1,
		Title: "Mehul",
		Author: entities.Author{
			Id:        1,
			FirstName: "Charles",
			LastName:  "Lee",
			Dob:       "12/01/1942",
			PenName:   "CL",
		},
		Publication:   "Penguin",
		PublishedDate: "12/02/1970",
		AuthorID:      1,
	}, nil
}

func (m mock) PostBook(books entities.Books) (int64, error) {
	return 0, nil
}

func (m mock) PutBook(books entities.Books, id int) (entities.Books, error) {
	return entities.Books{}, nil
}

func (m mock) DeleteBook(id int) (int64, error) {
	if id <= 0 {
		return 0, errors.New("invalid id")
	}
	return 1, nil
}
