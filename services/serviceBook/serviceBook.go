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

	//var books []entities.Book
	//var err error
	//fmt.Println("hello Mehul")
	//if title != "" {
	//	books, err = bs.bookstore.GetAllBooksByTitle(ctx, title)
	//	if err != nil {
	//		return []entities.Book{}, err
	//	} else {
	//		books, err = bs.bookstore.GetAllBooks(ctx, title, getAuthor)
	//		if err != nil {
	//			return []entities.Book{}, err
	//		}
	//	}
	//}
	//books, err = bs.bookstore.GetAllBooks(ctx, title, getAuthor)
	//if err != nil {
	//	//fmt.Println("hello get all service layer")
	//	return nil, nil
	//}

	//if getAuthor == "true" {
	//	for i := range books {
	//		res, err := bs.service
	//	}
	//}

	return []entities.Book{}, nil
}

///////get book by id  -- completed, running properly
func (bs serviceBook) GetBookByID(ctx context.Context, id int) (entities.Book, error) {
	if id <= 0 {
		return entities.Book{}, errors.New("negative id for book")
	}
	res, err := bs.bookstore.CheckBook(context.Background(), id)
	if err != nil {
		fmt.Println("hello")
		return entities.Book{}, err
	}
	if res == true {
		//chkauthor, err := bs.authorstore.CheckAuthor(context.Background(), )
		res, err := bs.bookstore.GetBookByID(context.Background(), id)
		if err != nil {
			return entities.Book{}, err
		}
		res.Author, err = bs.authorstore.GetAuthorByID(context.Background(), res.AuthorID)
		if err != nil {
			//fmt.Println("hi")
			return entities.Book{}, err
		}
		return res, nil
	}
	return entities.Book{}, errors.New("book does not exists")
}

////post book - completed
func (bs serviceBook) PostBook(ctx context.Context, books entities.Book) (int64, error) {
	chk, err := bs.authorstore.CheckAuthorByID(context.Background(), books.AuthorID)
	if err != nil {
		return 0, err
	}
	if chk == true {
		//	res, err := bs.bookstore.CheckBook(ctx, books.Id)
		//if err != nil {
		//	return 0, err
		//}
		//if res == true {
		res, err := bs.bookstore.PostBook(context.Background(), &books)
		if err != nil {
			return 0, err
		}
		return res, nil
		//}
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
		res, err := bs.bookstore.PutBook(ctx, book, id)
		if err != nil {

		}
		return res, nil
	}
	return entities.Book{}, errors.New("book id does not exists")
}

// delete book - completed
func (bs serviceBook) DeleteBook(ctx context.Context, id int) (int64, error) {
	if id <= 0 {
		return 0, errors.New("id is negative")
	}
	chk, err := bs.bookstore.CheckBook(context.Background(), id)
	if err != nil {
		return 0, err
	}
	if chk == true {
		res, err := bs.bookstore.DeleteBook(ctx, id)
		if err != nil {
			return 0, err
		}
		return res, nil // res is rows affected
	}
	return 0, err
}
