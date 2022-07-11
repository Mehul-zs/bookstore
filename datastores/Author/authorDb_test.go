package datastoreauthor

import (
	"Bookstore/entities"
	"context"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"net/http"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

//
//func DBConn() *sql.DB {
//	db, err := sql.Open("mysql", "root"+":"+"HelloMehul1@"+"@tcp(localhost:3306)"+"/"+"bookstore")
//	if err != nil {
//		log.Fatal("failed to connect with database:\n", err)
//	}
//
//	pingErr := db.Ping()
//	if pingErr != nil {
//		log.Fatal("failed to ping", pingErr)
//	}
//
//	return db
//}

/// get all author -- completed
func TestGetAllAuthor(t *testing.T) {
	testcases := []struct {
		desc    string
		expRows *sqlmock.Rows
		expRes  []entities.Author
		expErr  error
	}{
		{
			desc: "Get All Books", expRows: sqlmock.NewRows([]string{"id", "first_name", "last_name", "dob", "pen_name"}).
				AddRow(1, "Jp", "Morgan", "12/01/1999", "Morgan"),
			expRes: []entities.Author{{Id: 1, FirstName: "Jp", LastName: "Morgan", Dob: "12/10/1999", PenName: "Morgan"}},
		},
		{
			desc: "Error ", expRows: sqlmock.NewRows([]string{"id", "first_name", "last_name", "dob", "pen_name"}).
				AddRow("1", "RD", "Sharma", "2/11/1989", "Sharma"),
			expErr: fmt.Errorf("query err"),
		},
	}
	for i, tc := range testcases {

		DB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		authorStore := New(DB)

		mock.ExpectQuery("SELECT * FROM Author").WillReturnRows(tc.expRows).WillReturnError(tc.expErr)
		resp, err := authorStore.GetAllAuthor(context.Background())

		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, err, tc.expErr)
		}

		if !reflect.DeepEqual(resp, tc.expRes) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, tc.expRes)
		}
	}
}

// post author -- completed
func TestPostAuthor(t *testing.T) {

	testcases := []struct {
		desc   string
		req    entities.Author
		expOut int64
		err    error
	}{
		{"Valid Author", entities.Author{Id: 10, FirstName: "Mehul", LastName: "Rawal", Dob: "18/07/2000", PenName: "Me"}, http.StatusCreated, nil},
		{desc: "Author Already exists", req: entities.Author{1, "Hey", "kumar", "01/07/2000", "Me"}, expOut: http.StatusBadRequest, err: errors.New("Author Alreadyexists")},
	}

	//DB := DBConn()
	DB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error during the opening of database:%v\n", err)
	}
	defer DB.Close()

	for i, tc := range testcases {

		//rows := sqlmock.NewRows([]string{"Id"}).AddRow(tc.req.Id)

		//mock.ExpectQuery("SELECT Id FROM Author WHERE FirstName=? AND LastName=? AND  Dob=? AND PenName=? AND Id =?").WithArgs(tc.req.FirstName, tc.req.LastName, tc.req.Dob, tc.req.PenName, tc.req.Id).WillReturnRows(rows).WillReturnError(tc.err)
		mock.ExpectExec("INSERT INTO Author (Id, FirstName,LastName, Dob, PenName) VALUES (?, ? , ?, ?, ?)").WithArgs(tc.req.Id, tc.req.FirstName, tc.req.LastName, tc.req.Dob, tc.req.PenName).WillReturnResult(sqlmock.NewResult(0, tc.expOut))

		a := New(DB)
		res, err := a.PostAuthor(context.Background(), tc.req)
		if err != tc.err && res != tc.expOut {
			t.Errorf("Test %d case failed", i)
		}

	}
}

// put author -- completed
func TestPutAuthor(t *testing.T) {
	testcases := []struct {
		desc   string
		id     int
		input  entities.Author
		expOut entities.Author
		err    error
	}{
		{"Updated successfully", 1, entities.Author{Id: 1, FirstName: "ATUL", LastName: "kumar", Dob: "01/07/2000", PenName: "Me"},
			entities.Author{Id: 1, FirstName: "Mehul", LastName: "kumar", Dob: "01/07/2000", PenName: "Me"}, nil},
		{"ID does not exist", 9, entities.Author{Id: 9, FirstName: "Akshit", LastName: "Gupta", Dob: "06/07/2000", PenName: "rk"},
			entities.Author{}, nil},
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a database connection", err)
	}
	defer db.Close()
	for _, tc := range testcases {

		mock.ExpectExec("UPDATE Author SET FirstName = ?, LastName = ?, Dob = ? , PenName = ?, Id=?  WHERE Id =?").WithArgs(
			tc.input.FirstName, tc.input.LastName, tc.input.Dob, tc.input.PenName, tc.input.Id, tc.id).WillReturnError(tc.err)
		//fmt.Println("ME")
		a := New(db)
		//fmt.Println("Hii")

		res, err := a.PutAuthor(context.Background(), tc.input, tc.id)
		fmt.Println("Hello")

		if tc.err != err && res != tc.expOut {
			t.Errorf("Test case failed")
		}

	}

}

// delete author completed
func TestDeleteAuthor(t *testing.T) {
	testcases := []struct {
		desc   string
		input  int
		expOut int64
		err    error
	}{
		{"success:deleted", 3, 1, nil},
		{"failure:id does not exist", 5, 0, nil},
	}

	for i, tc := range testcases {

		DB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		authorStore := New(DB)

		mock.ExpectExec("Delete from Author WHERE Id=?").WithArgs(tc.input).WillReturnResult(sqlmock.NewResult(0, tc.expOut))
		res, err := authorStore.DeleteAuthor(context.TODO(), tc.input)

		if res != tc.expOut && tc.err != err {
			t.Errorf("failed %d for %v\n", i, tc.desc)
		}
	}

}
