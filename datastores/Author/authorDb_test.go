package datastoreAuthor

import (
	"Bookstore/entities"
	"database/sql"
	"database/sql/driver"
	"os"
	"reflect"
	"testing"
)

func initializeMySQL(t *testing.T) *sql.DB {
	conf := driver.MySQLConfig{
		Host:     os.Getenv("SQL_HOST"),
		User:     os.Getenv("SQL_USER"),
		Password: os.Getenv("SQL_PASSWORD"),
		Port:     os.Getenv("SQL_PORT"),
		Db:       os.Getenv("SQL_DB"),
	}

	var err error
	db, err := driver.ConnectToMySQL(conf)
	if err != nil {
		t.Errorf("could not connect to sql, err:%v", err)
	}

	return db
}

func TestDatastore(t *testing.T) {
	db := initializeMySQL(t)
	a := New(db)
	testAuthorstore_PostAuthor(t, a)
}

func testAuthorstore_PostAuthor(t *testing.T) {

	testcases := []struct {
		req      entities.Author
		response entities.Author
	}{
		{entities.Author{FirstName: "Mehul", LastName: "Rawal", Dob: "18/07/2000", PenName: "Me"}, entities.Author{3, "Mehul", "Rawal", "18/07/2000", "Me"}},
	}
	for i, v := range testcases {
		resp, _ := db.Create(v.req)

		if !reflect.DeepEqual(resp, v.response) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.response)
		}
	}
}
