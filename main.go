package Bookstore

import (
	datastoreAuthor "Bookstore/datastores/Author"
	datastoreBook "Bookstore/datastores/book"
	handlerAuthor "Bookstore/handlers/Author"
	handlerBook "Bookstore/handlers/books"
	"Bookstore/services/serviceAuthor"
	"Bookstore/services/serviceBook"
	"database/sql"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func DbConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root"+":"+"HelloMehul1@"+"@tcp(localhost:3306)"+"/"+"BookStore")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	r := mux.NewRouter()

	db, err := DbConn()
	if err != nil {
		return
	}

	bookstore := datastoreBook.New(db)
	servicebook := serviceBook.New(bookstore)
	handlerbook := handlerBook.New(servicebook)
	http.HandleFunc("/book", handlerbook.Handler)

	////fmt.Println("Hey Mehul!")

	authorstore := datastoreAuthor.New(db)
	serviceauth := serviceAuthor.New(authorstore)
	handlerauthor := handlerAuthor.New(serviceauth)
	http.HandleFunc("/author", handlerauthor.Handler)

	err = http.ListenAndServe(":5000", r)
	if err != nil {
		return
	}
}
