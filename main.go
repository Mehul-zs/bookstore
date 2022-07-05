package main

import (
	datastoreAuthor "Bookstore/datastores/Author"
	datastoreBook "Bookstore/datastores/book"
	handlerAuthor "Bookstore/handlers/Author"
	handlerBook "Bookstore/handlers/books"
	"Bookstore/services/serviceAuthor"
	"Bookstore/services/serviceBook"
	"database/sql"
	"fmt"
	"log"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func DbConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root"+":"+"HelloMehul1@"+"@tcp(localhost:3306)"+"/"+"bookstore")
	if err != nil {
		return nil, err
	}

	chk := db.Ping()
	if chk != nil {
		fmt.Println("Error, connection not established")
		return nil, nil
	}
	fmt.Println("Connection establsied")

	return db, nil
}

func main() {

	db, err := DbConn()
	if err != nil {
		return
	}

	bookstore := datastoreBook.New(db)
	servicebook := serviceBook.New(bookstore)
	handlerbook := handlerBook.New(servicebook)

	////fmt.Println("Hey Mehul!")

	authorstore := datastoreAuthor.New(db)
	serviceauth := serviceAuthor.New(authorstore)
	handlerauthor := handlerAuthor.New(serviceauth)
	fmt.Println("Hello main")

	r := mux.NewRouter()
	fmt.Println(db)

	r.HandleFunc("/books", handlerbook.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", handlerbook.GetByID).Methods(http.MethodGet)

	r.HandleFunc("/author", handlerauthor.PostAuthor).Methods(http.MethodPost)
	r.HandleFunc("/book", handlerbook.PostBook).Methods(http.MethodPost)

	r.HandleFunc("/author/{id}", handlerauthor.PutAuthor).Methods(http.MethodPut)
	r.HandleFunc("/books/{id}", handlerbook.PutBook).Methods(http.MethodPut)
	//
	r.HandleFunc("/deleteBook/{id}", handlerbook.DeleteBook).Methods(http.MethodDelete)
	r.HandleFunc("/deleteAuthor/{id}", handlerauthor.DeleteAuthor).Methods(http.MethodDelete)
	//fmt.Println("Hey Mehul!")

	Server := http.Server{
		Addr:    ":5000",
		Handler: r,
	}

	fmt.Println("Server started at 5000")
	log.Fatal(Server.ListenAndServe())
}
