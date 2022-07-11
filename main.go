package main

import (
	datastoreAuthor "Bookstore/datastore/Author"
	datastoreBook "Bookstore/datastore/book"
	handlerAuthor "Bookstore/handlers/Author"
	handlerBook "Bookstore/handlers/books"
	"Bookstore/services/serviceBook"
	"Bookstore/services/serviceauthor"
	"database/sql"
	"fmt"
	"log"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//connection to the database
func DBConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root"+":"+"HelloMehul1@"+"@tcp(localhost:3306)"+"/"+"bookstore")
	if err != nil {
		return nil, err
	}

	chk := db.Ping()
	if chk != nil {
		fmt.Println("Error, connection not established")
		return nil, nil
	}
	//fmt.Println("Connection establsied")

	return db, nil
}

func main() {

	db, err := DBConn()
	if err != nil {
		return
	}

	//mapping three layer architecture
	bookstore := datastoreBook.New(db)
	servicebook := serviceBook.New(bookstore)
	handlerbook := handlerBook.New(servicebook)

	authorstore := datastoreAuthor.New(db)
	serviceauth := serviceauthor.New(authorstore)
	handlerauthor := handlerAuthor.New(serviceauth)
	//fmt.Println("Hello main")

	//initialising the route
	r := mux.NewRouter()
	//fmt.Println(db)

	//r.HandleFunc("/books", handlerbook.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", handlerbook.GetBookByID).Methods(http.MethodGet)

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
	//starting the server at port 500
	fmt.Println("Server started at 5000")
	log.Fatal(Server.ListenAndServe())
}
