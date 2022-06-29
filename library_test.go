package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBook(t *testing.T) {
	testcases := []struct {
		desc   string
		input  *http.Request
		output []Books
	}{
		{"get books", httptest.NewRequest(http.MethodGet, "localhost:8000/book", nil), []Books{{1, "Godin", Author{}, "Penguin", "11/02/1988"}}},
	}

	for i, tc := range testcases {
		w := httptest.NewRecorder()
		GetBook(w, tc.input)
		//res := w.Result()

		res, _ := io.ReadAll(w.Result().Body)
		resBooks := []Books{}
		err := json.Unmarshal(res, &resBooks)
		if err != nil {
			return
		}

		for p := 1; p < len(resBooks); p++ {
			if resBooks[i] != tc.output[i] {
				t.Errorf("Error test case failed")
			}
		}

	}
}

func TestGetBookByID(t *testing.T) {
	testcasesId := []struct {
		desc    string
		input2  string
		output2 Books
	}{
		{"get books", "localhost:8000/book/1", Books{1, "Godin", Author{Id: 1}, "Penguin", "12/06/1899"}},
	}

	for _, tc := range testcasesId {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, tc.input2, nil)

		GetBook(w, req)
		//res := w.Result()

		res, _ := io.ReadAll(w.Result().Body)
		resBooks := Books{}
		err := json.Unmarshal(res, &resBooks)

		if err != nil {
			return
		}

		if resBooks != tc.output2 {
			t.Errorf("test caase failed in get by id case")
		}
		//fmt.Println(b)
	}

}

func TestPostByBook(t *testing.T) {
	testcases := []struct {
		desc       string
		postinput  Books
		postoutput Books
		status     int
	}{
		{"Details", Books{1, "linchpin", Author{}, "Penguin", "18/07/2000"}, Books{3, "Linchpin", Author{}, "Penguin", "18/07/2002"}, 200},
		{desc: "Invalid publication", postinput: Books{Id: 2, Title: "RD", Author: Author{}, Publication: "NA", PublishedDate: "12/03/2002"}, postoutput: Books{}, status: http.StatusBadRequest},
	}

	for i, tc := range testcases {
		rw := httptest.NewRecorder()
		body, _ := json.Marshal(tc)
		req := httptest.NewRequest(http.MethodPost, "localhost:8000/Book/", bytes.NewReader(body))

		PostByBook(rw, req)

		//defer rw.Result().Body.Close()

		if rw.Result().StatusCode != tc.status {
			t.Errorf("%v test case failed at %v, with status : %v", i, tc.desc, tc.status)
		}

		res, _ := io.ReadAll(rw.Result().Body)
		resBook := Books{}
		json.Unmarshal(res, &resBook)

		if resBook != tc.postoutput {
			t.Errorf("test case failed at %v : %v", i, tc.desc)
		}

	}

}

//func TestPostByAuthor(t *testing.T) {
//	testcases := []struct {
//		desc       string
//		postinput  Author
//		postoutput Author
//		status     int
//	}{
//		{"Valid details", Author{FirstName: "RD", LastName: "Sharma", Dob: "2/11/1989", PenName: "Sharma"}, Author{1, "RD", "Sharma", "2/11/1989", "Sharma"}, http.StatusOK},
//		{"InValid details", Author{FirstName: "", LastName: "Sharma", Dob: "2/11/1989", PenName: "Sharma"}, Author{}, http.StatusBadRequest},
//		{"Author already exists", Author{FirstName: "RD", LastName: "Sharma", Dob: "2/11/1989", PenName: "Sharma"}, Author{}, http.StatusBadRequest},
//	}
//	for i, tc := range testcases {
//		rw := httptest.NewRecorder()
//		body, _ := json.Marshal(tc.postinput)
//		req := httptest.NewRequest(http.MethodPost, "localhost:8000/Book/", bytes.NewReader(body))
//
//		if rw.Result().StatusCode != tc.status {
//			t.Errorf("%v test case failed at %v, with status : %v", i, tc.desc, tc.status)
//		}
//
//		res, _ := io.ReadAll(rw.Result().Body)
//		resBook := Author{}
//		json.Unmarshal(res, &resBook)
//
//		if resBook != tc.postoutput {
//			t.Errorf("test case failed at %v : %v", i, tc.desc)
//		}
//
//	}
//
//}

//func TestPutBookById(t *testing.T) {
//	testcases := []struct {
//		desc      string
//		id        int
//		InputById Books
//		status    int
//	}{{"Details", 1, Books{Title: "Linchpin", Author: Author{}, Publication: "Pengiun", PublishedDate: "11/03/2002"}, 200},
//		{"Invalid Publication", 1, Books{Title: "Rakshit", Author: Author{}, Publication: "Mehul", PublishedDate: "11/03/2002"}, http.StatusBadRequest},
//		{"Published date should be between 1880 and 2022", 1, Books{Title: "Check", Author: Author{}, Publication: "", PublishedDate: "1/1/1870"}, http.StatusBadRequest},
//		{"Published date should be between 1880 and 2022", 1, Books{Title: "James", Author: Author{}, Publication: "", PublishedDate: "1/1/2222"}, http.StatusBadRequest},
//		{"Author should exist", 1, Books{}, http.StatusBadRequest},
//		{"Title can't be empty", 1, Books{Title: "", Author: Author{}, Publication: "", PublishedDate: ""}, http.StatusBadRequest},
//	}
//
//	for _, tc := range testcases {
//
//		b, err := json.Marshal(tc.InputById)
//		if err != nil {
//			fmt.Println("error:", err)
//		}
//
//		req := httptest.NewRequest(http.MethodPost, "localhost:8000/book", bytes.NewBuffer(b))
//		w := httptest.NewRecorder()
//		PostByBook(w, req)
//
//		res := w.Result()
//		if res.StatusCode != tc.status {
//			t.Errorf("failed for %v\n", tc.desc)
//		}
//	}
//
//}

//func TestPutBook(t *testing.T) {
//	testcases := []struct {
//		desc        string
//		inputMethod string
//		target      string
//		body        Books
//		expected    int
//	}{}
//
//	for _, tc := range testcases {
//
//		ReqBody, err := json.Marshal(tc.body)
//		if err != nil {
//			fmt.Errorf("failed %v\n", err)
//		}
//		req := httptest.NewRequest(tc.inputMethod, "https://localhost:8000/book/{id}"+tc.target, bytes.NewBuffer(ReqBody))
//		req = mux.SetURLVars(req, map[string]string{"id": tc.target})
//		w := httptest.NewRecorder()
//		PutBook(w, req)
//
//		res := w.Result()
//		if res.StatusCode != tc.expected {
//			t.Errorf("failed for %s", tc.desc)
//		}
//		//assert.Equal(t, tc.expected, res.StatusCode)
//	}
//}

func TestPutAuthor(t *testing.T) {
	testcases := []struct {
		desc     string
		target   string
		author   Author
		expected int
	}{
		{"Id Already exists", "2", Author{2, "Mehul", "Rawal", "30/04/2001", "me"}, http.StatusCreated},
		{"Id Already exists", "1", Author{1, "Mehul", "Gupta", "30/04/2001", "me"}, http.StatusCreated},
		{"Id Inserted", "10", Author{10, "stephen", "hawkins", "30/04/2001", "SH"}, http.StatusBadRequest},
	}

	for _, tc := range testcases {

		ReqBody, err := json.Marshal(tc.author)
		if err != nil {
			fmt.Errorf("failed %v\n", err)
		}
		req := httptest.NewRequest(http.MethodPut, "https://localhost:8000/author/{id}"+tc.target, bytes.NewBuffer(ReqBody))
		req = mux.SetURLVars(req, map[string]string{"id": tc.target})
		w := httptest.NewRecorder()
		PutAuthor(w, req)

		res := w.Result()
		if res.StatusCode != tc.expected {
			t.Errorf("failed for %s", tc.desc)
			return
		}
		if tc.expected != res.StatusCode {
			t.Errorf("test case failed")
		}
	}
}

func TestDeleteBook(t *testing.T) {
	testcases := []struct {
		desc        string
		inputMethod string
		target      string
		expected    int
	}{
		{"valid id", "DELETE", "1", http.StatusNoContent},
		{"invalid id", "DELETE", "-1", http.StatusBadRequest},
	}

	for _, tc := range testcases {

		req := httptest.NewRequest(tc.inputMethod, "localhost:8000/deleteBook/{id}"+tc.target, nil)
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": tc.target})
		DeleteBook(w, req)

		res := w.Result()
		if res.StatusCode != tc.expected {
			t.Errorf("failed for %s", tc.desc)
		}

	}
}

func TestDeleteAuthor(t *testing.T) {
	testcases := []struct {
		desc        string
		inputMethod string
		target      string
		expected    int
	}{
		{"valid authorId", "DELETE", "2", http.StatusNoContent},
		//{"invalid authorId", "DELETE", "-2", http.StatusBadRequest},
	}

	for _, tc := range testcases {

		req := httptest.NewRequest(tc.inputMethod, "localhost:8000/deleteAuthor/{id}"+tc.target, nil)
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": tc.target})
		DeleteAuthor(w, req)

		res := w.Result()
		if res.StatusCode != tc.expected {
			t.Errorf("failed for %s", tc.desc)
		}

	}
}
