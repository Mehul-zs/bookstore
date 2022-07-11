package serviceBook

import (
	datastore "Bookstore/datastore"
	"Bookstore/entities"
	"context"
	"errors"
	"fmt"
)

type serviceBook struct {
	bookstore   datastore.BookStore
	authorstore datastore.AuthorStore
}

func New(b datastore.BookStore, a datastore.AuthorStore) serviceBook {
	return serviceBook{b, a}
}

func (bs serviceBook) GetAllBooks(ctx context.Context, title string, getAuthor string) ([]entities.Book, error) {

	var books []entities.Book
	var err error
	//fmt.Println("hello Mehul")
	if title != "" {
		books, err = bs.bookstore.GetAllBooksByTitle(ctx, title)
		if err != nil {
			return []entities.Book{}, err
		} else {
			books, err = bs.bookstore.GetAllBooks(ctx, title, getAuthor)
			if err != nil {
				return []entities.Book{}, err
			}
		}
	}
	books, err = bs.bookstore.GetAllBooks(ctx, title, getAuthor)
	if err != nil {
		//fmt.Println("hello get all service layer")
		return nil, nil
	}

	if getAuthor == "true" {
		for i := range books {
			res, err := bs.service
		}
	}

	return books, nil
}

/////// get book by id  -- completed, running properly
func (bs serviceBook) GetBookByID(ctx context.Context, id int) (entities.Book, error) {
	res, err := bs.bookstore.CheckBook(ctx, id)
	if err != nil {
		return entities.Book{}, err
	}
	if res {
		return bs.bookstore.GetBookByID(ctx, id)
	}
	return entities.Book{}, errors.New("book does not exists")
}

//// post book - completed
func (bs serviceBook) PostBook(ctx context.Context, books entities.Book) (int64, error) {
	res, err := bs.bookstore.CheckBook(ctx, books.Id)
	if err != nil {
		fmt.Println("Hello post book")
		return 0, err
	}
	if res {
		return bs.bookstore.PostBook(ctx, &books)
	}
	return 0, errors.New("book Already exists")

}

//// put book completed
func (bs serviceBook) PutBook(ctx context.Context, book entities.Book, id int) (entities.Book, error) {
	res, err := bs.bookstore.CheckBook(ctx, id)
	if err != nil {
		return entities.Book{}, err
	}
	if res == true {
		return bs.bookstore.PutBook(ctx, book, id)
	}
	return entities.Book{}, errors.New("book id does not exists")

}

// delete book - completed
func (bs serviceBook) DeleteBook(ctx context.Context, id int) (int64, error) {
	res, err := bs.bookstore.DeleteBook(ctx, id)
	if err != nil {
		return 0, err
	}
	return res, nil
}
