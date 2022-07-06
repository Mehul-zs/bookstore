package datastorebook

import (
	"Bookstore/entities"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type Bookstore struct {
	db *sql.DB
}

func New(db *sql.DB) Bookstore {
	return Bookstore{db: db}
}

func (b Bookstore) GetAll(title string, getauthor string) ([]entities.Books, error) {

	var rows *sql.Rows
	//if title == "" {
	rows, err := b.db.Query("select * from Books;")
	//} else {
	//	rows, err := b.db.Query("select * from Books where Title=?;", title)
	//}

	if err != nil {
		log.Print(err)
		return nil, nil
	}
	var books []entities.Books
	for rows.Next() {
		book := entities.Books{}
		err := rows.Scan(&book.Id, &book.Title, &book.Publication, &book.PublishedDate, &book.Author.Id)
		if err != nil {
			log.Print(err)
		}
		//fmt.Println("Hi")
		//if getAuthor == "true" {
		fmt.Println("Hello")
		row := b.db.QueryRow("SELECT * from Author WHERE Id=? ", book.Author.Id)
		err = row.Scan(&book.Author.Id, &book.Author.FirstName, &book.Author.LastName, &book.Author.Dob, &book.Author.PenName)
		if err != nil {
			//fmt.Println("GG")
			log.Fatal(err)
		}
		//}
		books = append(books, book)
	}

	return books, nil
}

func (b Bookstore) GetByID(ID int) (entities.Books, error) {
	bookrow := b.db.QueryRow("select * from Books where Id=?;", ID)
	book := entities.Books{}
	err := bookrow.Scan(&book.Id, &book.Title, &book.Publication, &book.PublishedDate, &book.Author.Id)

	if err != nil {
		return entities.Books{}, nil
	}
	authrow := b.db.QueryRow("select * from Author where Id=?;", book.Author.Id)
	err = authrow.Scan(&book.Author.Id, &book.Author.FirstName, &book.Author.LastName, &book.Author.Dob, &book.Author.PenName)
	if err != nil {
		fmt.Println("Error")
		log.Print(err)
	}
	//json.NewEncoder(rw).Encode(book)

	return entities.Books{}, nil
}

func (b Bookstore) PostBook(books entities.Books) (int64, error) {
	result, err := b.db.Query("SELECT Id FROM Author WHERE FirstName = ? AND LastName =? AND  Dob = ? AND PenName =?",
		books.Author.FirstName, books.Author.LastName, books.Author.Dob, books.Author.PenName)
	if err != nil {
		fmt.Println("Author does not exists")
		return http.StatusBadRequest, nil
	}

	var authId int64
	if result.Next() {
		err = result.Scan(&authId)
		if err != nil {
			log.Print("author is not present")
		}
	} else {
		res, err := b.db.Exec("INSERT INTO Author (Id, FirstName,LastName,Dob,PenName) VALUES (?, ?, ?, ?, ?)",
			books.Author.Id, books.Author.FirstName, books.Author.LastName, books.Author.Dob, books.Author.PenName)
		if err != nil {
			return http.StatusBadRequest, nil
		}
		_, err = res.LastInsertId()
		if err != nil {
			return http.StatusBadRequest, nil
		}
	}
	//fmt.Println("last, Hello Mehul")

	if !BookExists(books, authId, b) {
		_, err = b.db.Exec("INSERT INTO Books (Id, Title, Publication,PublishedDate, AuthorId) VALUES (? ,?, ?, ?, ?)",
			books.Id, books.Title, books.Publication, books.PublishedDate, authId)
		if err != nil {
			fmt.Println(err)
			return http.StatusBadRequest, nil
		}
	}
	return http.StatusCreated, nil
}

func (b Bookstore) PutBook(book entities.Books, id int) (entities.Books, error) {
	result, err := b.db.Query("SELECT Id FROM Author WHERE Id = ?", book.Author.Id)
	if err != nil {
		fmt.Println("Hello put function datastore")
		//rw.WriteHeader(http.StatusBadRequest)
		return entities.Books{}, nil
	}

	for !result.Next() {
		log.Print("author not present", book.Author.Id)
		return entities.Books{}, nil
	}

	result, err = b.db.Query("SELECT * FROM Books WHERE Id = ?")
	if err != nil {
		log.Print(err)
	}
	for !result.Next() {
		log.Print("Book not present")
		//rw.WriteHeader(http.StatusBadRequest)
		return entities.Books{}, nil
	}

	_, err = b.db.Exec("UPDATE Books SET Id=?, Title = ? ,Publication = ? ,PublishedDate=?  WHERE Id =?",
		book.Id, book.Title, book.Publication, book.PublishedDate, id)
	if err != nil {
		//rw.WriteHeader(http.StatusBadRequest)
		return entities.Books{}, nil
	}

	return book, nil
}

func (b Bookstore) DeleteBook(ID int) (int64, error) {

	_, err := b.db.Exec("DELETE from Books WHERE Id=?", ID)
	if err != nil {
		return http.StatusBadRequest, nil
	}

	return http.StatusNoContent, nil
}

func BookExists(book entities.Books, authId int64, b Bookstore) bool {
	res, err := b.db.Query("SELECT Id FROM Book WHERE Title = ? AND Publication =? AND  PublsihedDate= ? AND AuthorId =?",
		book.Title, book.Publication, book.PublishedDate, authId)
	if err != nil || !res.Next() {
		return false
	}
	return true
}
