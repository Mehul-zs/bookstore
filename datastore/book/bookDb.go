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

func (b Bookstore) GetAllBooks(ctx context.Context, title, getauthor string) ([]entities.Book, error) {
	return []entities.Book{}, nil
}

func (b Bookstore) GetBookByID(ctx context.Context, id int) (entities.Book, error) {
	bookrow := b.db.QueryRowContext(ctx, "select * from Books where Id=?", id)
	book := entities.Book{}
	err := bookrow.Scan(&book.Id, &book.Title, &book.Publication, &book.PublishedDate, &book.AuthorID)

	if err != nil {
		return entities.Book{}, nil
	}
	//authrow := b.db.QueryRow("select * from Author where Id=?", book.Author.Id)
	//err = authrow.Scan(&book.Author.Id, &book.Author.FirstName, &book.Author.LastName, &book.Author.Dob, &book.Author.PenName)
	//if err != nil {
	//	log.Print(err)
	//	return entities.Book{}, errors.New("invalid author id match")
	//}

	return book, nil
}

func (b Bookstore) PostBook(ctx context.Context, books *entities.Book) (int64, error) {
	res, err := b.db.Exec("INSERT INTO Books (Id, Title, Publication,PublishedDate, AuthorId) VALUES (? ,?, ?, ?, ?)",
		&books.Id, &books.Title, &books.Publication, &books.PublishedDate, &books.AuthorID)
	ans, err := res.RowsAffected()

	if err != nil {
		fmt.Println(err)
		return 0, nil
	}

	return ans, nil
}

func (b Bookstore) PutBook(ctx context.Context, book entities.Book, id int) (entities.Book, error) {
	_, err := b.db.Exec("UPDATE Books SET Id=?, Title=?, Publication=?, PublishedDate=?, AuthorId=? WHERE Id =?",
		&book.Id, &book.Title, &book.Publication, &book.PublishedDate, &book.AuthorID, id)
	if err != nil {
		return entities.Book{}, nil
	}

	return book, nil
}

func (b Bookstore) DeleteBook(ctx context.Context, id int) (int64, error) {
	row, err := b.db.Exec("DELETE from Books WHERE Id=?", id)
	cnt, _ := row.RowsAffected()
	if err != nil {
		return 0, errors.New("No rows affected, id does not exist")
	}
	return cnt, nil
}

func BookExists(book *entities.Book, authId int64, b Bookstore) bool {
	res, err := b.db.Query("SELECT Id FROM Books WHERE Title = ? AND Publication =? AND  PublsihedDate= ? AND AuthorId =?",
		book.Title, book.Publication, book.PublishedDate, authId)
	if err != nil || !res.Next() {
		return false
	}
	return true
}

func (bs Bookstore) CheckBook(ctx context.Context, id int) (bool, error) {
	_, err := bs.db.Query("select * from Books where Id=?", id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (bs Bookstore) GetAllBooksByTitle(ctx context.Context, title string) ([]entities.Book, error) {
	var (
		books []entities.Book
		Rows  *sql.Rows
		err   error
	)

	Rows, err = bs.db.Query("SELECT * FROM Books WHERE title=?", title)
	if err != nil {
		log.Print(err)
		return []entities.Book{}, err
	}
	defer Rows.Close()

	for Rows.Next() {
		var book entities.Book

		err = Rows.Scan(&book.Id, &book.Title, &book.Publication, &book.PublishedDate, &book.Author.Id)
		if err != nil {
			return []entities.Book{}, err
		}

		books = append(books, book)
	}

	return books, nil
}
