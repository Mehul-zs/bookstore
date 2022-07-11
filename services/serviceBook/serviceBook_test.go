package serviceBook

import (
	"Bookstore/datastore"
	"Bookstore/entities"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"log"
	"net/http"
	"testing"
)

//func TestServiceBook_GetAll(t *testing.T) {
//	testcases := []struct {
//		input  entities.Books
//		output []entities.Books
//	}{
//		{},
//	}
//
//	for _, tc := range testcases {
//		b := New(mockDatastore{})
//		res, err := b.GetAll("", "")
//	}
//
//}

func TestServiceBook_GetByID(t *testing.T) {
	testcases := []struct {
		desc   string
		input  int
		expOut entities.Book
		err    error
	}{
		{desc: "Valid book details", input: 1, expOut: entities.Book{Id: 1, Title: "james", Author: entities.Author{Id: 1, FirstName: "Mehul", LastName: "Rawal", Dob: "18/07/2000", PenName: "Me"}, Publication: "Penguin", PublishedDate: "12/01/2020", AuthorID: 1}, err: nil},
		{"Invalid book details", -1, entities.Book{}, errors.New("Negative book id")},
	}

	for i, tc := range testcases {
		b := New(mockDatastore{})
		res, err := b.GetByID(tc.input)
		if err != tc.err && tc.expOut != res {
			t.Errorf("testcase:%d desc :%s actualError:%v actualResult:%v expectedError:%v expectedResult:%v", i, tc.desc, err, res, tc.err, tc.expOut)
		}

	}
}

func TestServiceBook_PostBook(t *testing.T) {

	testcases := []struct {
		desc      string
		req       entities.Book
		expStatus int64
		err       error
	}{
		{"Valid book", entities.Book{Id: 1, Title: "james", Author: entities.Author{Id: 1, FirstName: "Mehul", LastName: "Rawal", Dob: "18/07/2000", PenName: "Me"}, Publication: "Penguin", PublishedDate: "12/01/2020", AuthorID: 1}, http.StatusCreated, nil},
		{"Invalid published date", entities.Book{Id: 3, Title: "John Gems", Author: entities.Author{FirstName: "Mehul", LastName: "Rawal", Dob: "18/07/2000", PenName: "Me"}, Publication: "Penguin", PublishedDate: "12/02/2222"}, http.StatusBadRequest, errors.New("invalid published year")},
	}
	for i, tc := range testcases {
		b := New(mockDatastore{})
		res, err := b.PostBook(tc.req)
		if err != tc.err && tc.expStatus != res {
			t.Errorf("testcase:%d desc :%s actualError:%v actualResult:%v expectedError:%v expectedResult:%v", i, tc.desc, err, res, tc.err, tc.expStatus)

		}
	}

}

func TestServiceBook_PutBook(t *testing.T) {
	testcases := []struct {
		desc   string
		id     int
		input  entities.Book
		expOut entities.Book
		err    error
	}{
		{"Valid book", 1,
			entities.Book{Id: 3, Title: "Here we go", Author: entities.Author{Id: 5, FirstName: "Mehul", LastName: "Testing", Dob: "18/07/1960", PenName: "m"}, Publication: "Penguin", PublishedDate: "12/03/1980", AuthorID: 5},
			entities.Book{Id: 3, Title: "Here we go", Author: entities.Author{Id: 5, FirstName: "Mehul", LastName: "Testing", Dob: "18/07/1960", PenName: "m"}, Publication: "Penguin", PublishedDate: "12/03/1980", AuthorID: 5},
			nil},
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tc := range testcases {
		mockBookstore := datastore.NewMockBookStore(mockCtrl)
		mockAuthorstore := datastore.NewMockAuthorStore(mockCtrl)
		bs := New(mockBookstore, mockAuthorstore)

	}

}

func TestServiceBook_DeleteBook(t *testing.T) {
	testcases := []struct {
		desc   string
		input  int
		expOut int64
		err    error
	}{
		{"Valid BookId", 2, http.StatusNoContent, nil},
		{"Invalid BookId", -5, http.StatusBadRequest, errors.New("Negative ID params")},
		{"Valid BookId", 4, http.StatusBadRequest, errors.New("invalid book details")},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	for _, tc := range testcases {
		mockBookstore := datastore.NewMockBookStore(mockCtrl)
		mockAuthorstore := datastore.NewMockAuthorStore(mockCtrl)
		bs := New(mockBookstore, mockAuthorstore)
		mockBookstore.EXPECT().DeleteBook(context.Background(), tc.input).Return(tc.expOut, tc.err).AnyTimes()
		res, err := bs.DeleteBook(context.TODO(), tc.input)
		if err != nil {
			log.Println("book not deleted")
		}
		if tc.expOut != res {
			log.Println(tc.expOut, res)
		}
	}
}
