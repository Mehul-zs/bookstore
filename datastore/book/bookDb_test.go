package datastorebook

import (
	"Bookstore/entities"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"testing"
)

func TestBookstore_GetAllBooks(t *testing.T) {
	testcases := []struct {
		desc   string
		output []entities.Book
	}{
		{"Validated", []entities.Book{{
			Id:    1,
			Title: "James",
			Author: entities.Author{
				Id:        1,
				FirstName: "Mehul",
				LastName:  "Gupta",
				Dob:       "12/02/1970",
				PenName:   "Me",
			},
			Publication:   "Penguin",
			PublishedDate: "12/07/1999",
			AuthorID:      1,
		}},
		},
	}

	for _, tc := range testcases {
		DB := DBConn()
		bookStore := New(DB)

		chk, err := bookStore.GetAllBooks(_, _, _)

		if err != nil {
			t.Errorf("Test case failed %v", tc.desc)
		}

		if !reflect.DeepEqual(chk, tc.output) {
			t.Errorf("Test case failed: %s", tc.output)
		}
	}
}

func TestBookstore_GetBookByID(t *testing.T) {
	testcases := []struct {
		desc   string
		input  int
		expOut entities.Book
		err    error
	}{
		{"success", 1, entities.Book{Id: 1, AuthorID: 1, Title: "Clash of Titans", Publication: "penguin", PublishedDate: "16/07/1990", Author: entities.Author{Id: 1, FirstName: "Mehul", LastName: "Kumar", Dob: "01/04/1971", PenName: "Jonhy"}}, nil},
		{"Invalid book Id", -2, entities.Book{}, errors.New("invalid ID")},
	}

	DB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	b := New(DB)

	if err != nil {
		t.Fatalf("error during the opening of database:%v\n", err)
	}
	defer DB.Close()

	for _, tc := range testcases {
		rows := sqlmock.NewRows([]string{"Id", "Title", "Publication", "PublishedDate", "AuthorId"}).AddRow(tc.expOut.Id,
			tc.expOut.Title, tc.expOut.Publication, tc.expOut.PublishedDate, tc.expOut.AuthorID)
		authrows := sqlmock.NewRows([]string{"Id", "FirstName", "LastName", "Dob", "PenName"}).AddRow(tc.expOut.Author.Id,
			tc.expOut.Author.FirstName, tc.expOut.Author.LastName, tc.expOut.Author.Dob, tc.expOut.Author.PenName)
		mock.ExpectQuery("select * from Books where Id=?").WithArgs(tc.input).WillReturnRows(rows)
		mock.ExpectQuery("select * from Author where Id=?").WithArgs(tc.expOut.AuthorID).WillReturnRows(authrows)

		chk, err := b.GetBookByID(context.Background(), tc.input)

		if chk != tc.expOut && err != tc.err {
			t.Errorf("Test case failed; expout : %v, output: %v, err: %s, experr: %v", tc.expOut, chk, err, tc.err)
		}
	}

}

func TestBookstore_PostBook(t *testing.T) {
	testcases := []struct {
		desc           string
		req            *entities.Book
		lastInsertedId int64
		rowsAffected   int64
		err            error
	}{
		{"Valid book", &entities.Book{10, "james", entities.Author{Id: 2, FirstName: "Mehul", LastName: "Rawal", Dob: "18/07/2000", PenName: "Me"}, "Penguin", "12/01/2020", 2}, 0, 1, nil},
		{"Invalid published date", &entities.Book{10, "Rakshit", entities.Author{Id: 3, FirstName: "Mehul", LastName: "Rawal", Dob: "18/07/2000", PenName: "e"}, "Penguin", "12/02/2222", 3}, 0, 0, nil},
	}

	DB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	b := New(DB)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer DB.Close()

	for i, tc := range testcases {
		mock.ExpectExec("INSERT INTO Books (Id, Title, Publication,PublishedDate, AuthorId) VALUES (? ,?, ?, ?, ?)").WithArgs(tc.req.Id, tc.req.Title, tc.req.Publication, tc.req.PublishedDate, tc.req.AuthorID).WillReturnResult(sqlmock.NewResult(tc.lastInsertedId, tc.rowsAffected)).WillReturnResult(sqlmock.NewResult(0, tc.rowsAffected)).WillReturnError(tc.err)

		row, err := b.PostBook(context.Background(), tc.req)
		if err != tc.err && tc.rowsAffected != row {
			t.Errorf("Test %d case failed, err : %v, experr: %v, output: %v, expOut: %v", i, err, tc.err, row, tc.rowsAffected)
		}
	}
}

func TestBookstore_PutBook(t *testing.T) {
	testcases := []struct {
		desc         string
		id           int
		input        *entities.Book
		output       *entities.Book
		lastInsertId int64
		rowsafected  int64
		err          error
	}{
		{"Valid book", 10, &entities.Book{10, "james", entities.Author{Id: 2, FirstName: "Mehul", LastName: "Rawal", Dob: "18/07/2000", PenName: "Me"},
			"Penguin", "12/01/2020", 2}, &entities.Book{10, "james", entities.Author{Id: 2, FirstName: "Mehul", LastName: "Rawal", Dob: "18/07/2000", PenName: "Me"},
			"Penguin", "12/01/2020", 2}, 0, 1, nil},
	}
	DB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	b := New(DB)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	for i, tc := range testcases {
		mock.ExpectExec("UPDATE Books SET Id=?, Title=?, Publication=?, PublishedDate=?, AuthorId=? WHERE Id =?").WithArgs(tc.input.Id, tc.input.Title,
			tc.input.Publication, tc.input.PublishedDate, tc.input.AuthorID, tc.input.Id).WillReturnResult(sqlmock.NewResult(0, tc.rowsafected)).WillReturnError(tc.err)

		row, err := b.PutBook(context.Background(), tc.input, tc.id)
		if err != tc.err && tc.output != row {
			t.Errorf("Test %d case failed, err : %v, experr: %v, output: %v, expOut: %v", i, err, tc.err, row, tc.rowsafected)
		}
	}

}

func TestBookstore_DeleteBook(t *testing.T) {

	testcases := []struct {
		desc           string
		id             int
		rowsaffected   int64
		lastInsertedId int64
		err            error
	}{

		{"Deleted successfully", 2, 1, 0, nil},
		{"Delete Unsuccessful", 0, 0, 0, errors.New("No rows affected, id does not exist")},
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	for i, tc := range testcases {

		b := New(db)
		mock.ExpectExec("DELETE from Books WHERE Id=?").WithArgs(tc.id).WillReturnResult(sqlmock.NewResult(tc.lastInsertedId, tc.rowsaffected))

		row, err := b.DeleteBook(context.TODO(), tc.id)
		if err != tc.err && tc.rowsaffected != row {
			t.Errorf("Test %d case failed, err : %v, experr: %v", i, err, tc.err)
		}

	}

}
