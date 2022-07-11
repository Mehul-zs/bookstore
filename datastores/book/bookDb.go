package datastorebook

import (
	"Bookstore/entities"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Bookstore struct {
	db *sql.DB
}

func New(db *sql.DB) Bookstore {
	return Bookstore{db: db}
}

func (b Bookstore) GetAllBooks(title string, getauthor string) ([]entities.Book, error) {
	return []entities.Book{}, nil
}

// correct return error things..
func (b Bookstore) GetBookByID(ID int) (entities.Book, error) {
	bookrow := b.db.QueryRow("select * from Books where Id=?", ID)
	book := entities.Book{}
	err := bookrow.Scan(&book.Id, &book.Title, &book.Publication, &book.PublishedDate, &book.Author.Id)

	if err != nil {
		fmt.Println("Hello Mehul case failed")
		return entities.Book{}, nil
	}
	authrow := b.db.QueryRow("select * from Author where Id=?", book.Author.Id)
	err = authrow.Scan(&book.Author.Id, &book.Author.FirstName, &book.Author.LastName, &book.Author.Dob, &book.Author.PenName)
	if err != nil {
		fmt.Println("Hello Mehul welcome back to error")
		log.Print(err)
		return entities.Book{}, errors.New("invalid author id match")
	}
	//json.NewEncoder(rw).Encode(book)

	return book, nil
}

// post book -- completed
func (b Bookstore) PostBook(ctx context.Context, books entities.Book) (int64, error) {
	//result, err := b.db.Query("SELECT Id FROM Author WHERE FirstName = ? AND LastName =? AND  Dob = ? AND PenName =?",
	//	books.Author.FirstName, books.Author.LastName, books.Author.Dob, books.Author.PenName)
	//if err != nil {
	//	fmt.Println(err)
	//	fmt.Println("Author does not exists")
	//	return 0, nil
	//}

	//var authId int64

	//if !result.Next() {
	//	err = result.Scan(&authId)
	//	if err != nil {
	//		log.Print("author is not present")
	//		return 0, nil
	//	}
	//} else {
	//	res, err := b.db.Exec("INSERT INTO Author (Id, FirstName,LastName,Dob,PenName) VALUES (?, ?, ?, ?, ?)",
	//		books.Author.Id, books.Author.FirstName, books.Author.LastName, books.Author.Dob, books.Author.PenName)
	//	if err != nil {
	//		return 0, nil
	//	}
	//	ans, err = res.LastInsertId()
	//	if err != nil {
	//		return 0, nil
	//	}
	//}
	//fmt.Println("last, Hello Mehul")

	//if !BookExists(books, authId, b) {
	res, err := b.db.Exec("INSERT INTO Books (Id, Title, Publication,PublishedDate, AuthorId) VALUES (? ,?, ?, ?, ?)",
		books.Id, books.Title, books.Publication, books.PublishedDate, books.AuthorID)
	ans, err := res.RowsAffected()

	if err != nil {
		fmt.Println(err)
		return 0, nil
	}
	//}
	return ans, nil
}

// put book function working fine  -  completed
func (b Bookstore) PutBook(ctx context.Context, book entities.Book, id int) (entities.Book, error) {

	_, err := b.db.Exec("UPDATE Books SET Id=?, Title=?, Publication=?, PublishedDate=?, AuthorId=? WHERE Id =?",
		book.Id, book.Title, book.Publication, book.PublishedDate, book.AuthorID, id)
	if err != nil {
		return entities.Book{}, nil
	}

	return book, nil
}

//working fine - completed
func (b Bookstore) DeleteBook(ctx context.Context, ID int) (int64, error) {

	row, err := b.db.Exec("DELETE from Books WHERE Id=?", ID)
	cnt, _ := row.RowsAffected()
	if err != nil || ID == 0 {
		return 0, errors.New("No rows affected, id does not exist")
	}
	return cnt, nil
}

func BookExists(book entities.Book, authId int64, b Bookstore) bool {
	res, err := b.db.Query("SELECT Id FROM Books WHERE Title = ? AND Publication =? AND  PublsihedDate= ? AND AuthorId =?",
		book.Title, book.Publication, book.PublishedDate, authId)
	if err != nil || !res.Next() {
		return false
	}
	return true
}

func (bs Bookstore) CheckBook(ctx context.Context, id int) (bool, error) {
	row, err := bs.db.Query("select * from Books where BookId=?", id)
	if err != nil || !row.Next() {
		return false, err
	}

	return true, nil
}
