package datastorebook

import (
	"Bookstore/entities"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"testing"
)

func DBConn() *sql.DB {
	db, err := sql.Open("mysql", "root"+":"+"HelloMehul1@"+"@tcp(localhost:3306)"+"/"+"bookstore")
	if err != nil {
		log.Fatal("failed to connect with database:\n", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("failed to ping", pingErr)
	}

	return db
}

//func TestBookstore_GetAll(t *testing.T) {
//	testcases := []struct {
//		desc   string
//		output []entities.Books
//	}{
//		{"Validated", []entities.Books{{
//			Id:    1,
//			Title: "James",
//			Author: entities.Author{
//				Id:        1,
//				FirstName: "Mehul",
//				LastName:  "Gupta",
//				Dob:       "12/02/1970",
//				PenName:   "Me",
//			},
//			Publication:   "Penguin",
//			PublishedDate: "12/07/1999",
//			AuthorID:      1,
//		}},
//		},
//	}
//
//	for _, tc := range testcases {
//		DB := DBConn()
//		bookStore := New(DB)
//
//		chk, err := bookStore.GetAll()
//
//		if err != nil {
//			t.Errorf("Test case failed %v", tc.desc)
//		}
//
//		if !reflect.DeepEqual(chk, tc.output) {
//			t.Errorf("Test case failed: %s", tc.output)
//		}
//	}
//}

func TestBookstore_GetByID(t *testing.T) {
	testcases := []struct {
		desc   string
		input  int
		expOut entities.Books
		err    error
	}{
		{"success", 1, entities.Books{Id: 1, AuthorID: 1, Title: "Clash of Titans", Publication: "penguin", PublishedDate: "16/07/1990", Author: entities.Author{Id: 1, FirstName: "Mehul", LastName: "Kumar", Dob: "01/04/1971", PenName: "Jonhy"}}, nil},
		{"Invalid book Id", -2, entities.Books{}, errors.New("invalid ID")},
	}
	for _, tc := range testcases {

		DB := DBConn()

		bookStore := New(DB)

		chk, err := bookStore.GetByID(tc.input)

		if chk != tc.expOut {
			t.Errorf("Test case failed; expout : %v, output: %v, err: %s", tc.expOut, chk, err)
		}
	}

}

func TestBookstore_PostBook(t *testing.T) {
	testcases := []struct {
		desc      string
		req       entities.Books
		expStatus int64
		//err       error
	}{
		{"Valid book", entities.Books{10, "james", entities.Author{Id: 2, FirstName: "Mehul", LastName: "Rawal", Dob: "18/07/2000", PenName: "Me"}, "Penguin", "12/01/2020", 2}, http.StatusCreated},
		{"Invalid published date", entities.Books{1, "Rakshit", entities.Author{Id: 3, FirstName: "Mehul", LastName: "Rawal", Dob: "18/07/2000", PenName: "e"}, "Penguin", "12/02/2222", 3}, http.StatusBadRequest},
	}

	DB := DBConn()
	bookStore := New(DB)

	for i, tc := range testcases {

		res, _ := bookStore.PostBook(tc.req)
		if res != tc.expStatus {
			t.Errorf("testcase:%d desc:%v actualoutput:%v expectedoutput:%v", i, tc.desc, res, tc.expStatus)
		}
	}

}

func TestBookstore_DeleteBook(t *testing.T) {

	testcases := []struct {
		desc   string
		input  int
		expOut int64
		err    error
	}{

		{"Deleted successfully", 2, http.StatusNoContent, nil},
		{"ID does not exist", 0, http.StatusBadRequest, errors.New("Invalid ID")},
	}

	for _, tc := range testcases {

		DB := DBConn()
		bookstore := New(DB)

		id, err := bookstore.DeleteBook(tc.input)

		if tc.expOut != id && tc.err != err {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}

}
