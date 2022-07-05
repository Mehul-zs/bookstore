package serviceBook

import (
	datastore "Bookstore/datastores"
	"Bookstore/entities"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type serviceBook struct {
	bookstore datastore.BookStore
}

func New(b datastore.BookStore) serviceBook {
	return serviceBook{b}
}

func (bs serviceBook) GetAll(title string, getAuthor string) ([]entities.Books, error) {

	var books []entities.Books
	fmt.Println("hello Mehul")
	books, err := bs.bookstore.GetAll(title, getAuthor)
	if err != nil {
		fmt.Println("hello get all service layer")
		return nil, nil
	}

	return books, nil
}

func (bs serviceBook) GetByID(id int) (entities.Books, error) {
	if id <= 0 {
		return entities.Books{}, nil
	}
	var book entities.Books
	book, err := bs.bookstore.GetByID(id)
	if err != nil {
		return entities.Books{}, nil
	}
	return book, nil
}

func (bs serviceBook) PostBook(books entities.Books) (int64, error) {
	if books.Title == "" || books.Author.FirstName == "" {
		return http.StatusBadRequest, nil
	}

	if books.Publication != "Scholastic" && books.Publication != "Arihant" && books.Publication != "Penguin" {
		log.Print("Invalid Publication")
		return http.StatusBadRequest, nil
	}

	publishedDate := strings.Split(books.PublishedDate, "/")
	if len(publishedDate) < 3 {
		return http.StatusBadRequest, nil
	}
	yr, _ := strconv.Atoi(publishedDate[2])

	if yr >= 2022 || yr < 1880 {
		return http.StatusBadRequest, nil

	}

	_, err := bs.bookstore.PostBook(books)
	if err != nil {
		return http.StatusBadRequest, nil
	}

	return http.StatusCreated, nil

}

func (bs serviceBook) PutBook(book entities.Books, id int) (entities.Books, error) {
	fmt.Println("Hello Put book")
	if book.Title == "" {
		fmt.Println("Hello put title is empty")
		return entities.Books{}, nil
	}
	if book.Publication != "Scholastic" && book.Publication != "Arihant" && book.Publication != "Penguin" {
		log.Print("Invalid Publication")
		return entities.Books{}, nil
	}

	publicationDate := strings.Split(book.PublishedDate, "/")
	if len(publicationDate) < 3 {
		fmt.Println("Hello publication date")
		return entities.Books{}, nil
	}
	yr, _ := strconv.Atoi(publicationDate[2])
	if yr > time.Now().Year() || yr < 1880 {
		return entities.Books{}, nil
	}
	fmt.Println("Hello entering daatastore put book")
	books, err := bs.bookstore.PutBook(book, id)
	if err != nil {
		fmt.Println("Hello books")
		return entities.Books{}, nil
	}
	fmt.Println("Hello Mehul")
	return books, nil

}

func (bs serviceBook) DeleteBook(id int) (int64, error) {

	if id <= 0 {
		return http.StatusBadRequest, nil
	}

	cnt, err := bs.bookstore.DeleteBook(id)
	if err != nil {
		return http.StatusBadRequest, nil
	}

	return cnt, nil
}
