package datastoreauthor

import (
	"Bookstore/entities"
	"database/sql"
	"log"
	"net/http"
	"testing"

	_ "github.com/go-sql-driver/mysql"
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
func TestPostAuthor(t *testing.T) {

	testcases := []struct {
		desc   string
		req    entities.Author
		expOut int64
		//err    error
	}{
		{"Valid Author", entities.Author{Id: 2, FirstName: "Mehul", LastName: "Rawal", Dob: "18/07/2000", PenName: "Me"}, http.StatusCreated},
		{desc: "Author Already exists", req: entities.Author{1, "Rakshit", "Gupta", "12/02/1996", "kinu"}, expOut: http.StatusBadRequest},
		{"Valid Author", entities.Author{Id: 8, FirstName: "", LastName: "", Dob: "", PenName: ""}, http.StatusBadRequest},
	}

	DB := DBConn()
	authorStore := New(DB)

	for i, tc := range testcases {

		res, _ := authorStore.PostAuthor(tc.req)
		if res != tc.expOut {
			t.Errorf("testcase:%d desc:%v actualoutput:%v expectedoutput:%v", i, tc.desc, res, tc.expOut)
		}

	}
}

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

	for _, tc := range testcases {

		DB := DBConn()
		authorStore := New(DB)

		res, err := authorStore.DeleteAuthor(tc.input)

		if res != tc.expOut && tc.err != err {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestPutAuthor(t *testing.T) {
	testcases := []struct {
		desc   string
		id     int
		input  entities.Author
		expOut entities.Author
		err    error
	}{
		{"Updated successfully", 1, entities.Author{Id: 1, FirstName: "Mehul", LastName: "kumar", Dob: "01/07/2000", PenName: "Me"},
			entities.Author{Id: 1, FirstName: "Mehul", LastName: "kumar", Dob: "01/07/2000", PenName: "Me"}, nil},
		{"ID does not exist", 9, entities.Author{Id: 9, FirstName: "Rakshit", LastName: "Gupta", Dob: "06/07/2000", PenName: "rk"},
			entities.Author{}, nil},
	}
	db := DBConn()
	a := New(db)
	for i, tc := range testcases {
		res, err := a.PutAuthor(tc.input, tc.id)
		//assert.Equal(t, res, tc.response)
		//reflect.DeepEqual(tc.expected, res.StatusCode)
		if res != tc.expOut || err != tc.err {
			t.Errorf("testcase:%d desc:%v actualResult:%v actualError:%v expectedResponse:%v expectedError:%v", i, tc.desc, res, err, tc.expOut, tc.err)
		}
	}
}
