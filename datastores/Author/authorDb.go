package datastoreauthor

import (
	"Bookstore/entities"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type Authorstore struct {
	db *sql.DB
}

func New(db *sql.DB) Authorstore {
	return Authorstore{db: db}
}

// get all author completed
func (s Authorstore) GetAllAuthor(ctx context.Context) ([]entities.Author, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT * FROM Author")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	authors := make([]entities.Author, 0)

	for rows.Next() {
		var author entities.Author
		err = rows.Scan(author.Id, author.FirstName, author.LastName, author.Dob, author.PenName)
		if err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (a Authorstore) CheckAuthor(ctx context.Context, author entities.Author) (int, error) {

	_, err := a.db.QueryContext(ctx, "SELECT Id FROM Author WHERE FirstName=? AND LastName=? AND  Dob=? AND PenName=? AND Id =?", author.FirstName, author.LastName, author.Dob, author.PenName, author.Id)
	if err != nil {
		return 0, errors.New("Author Alreadyexists")

	}

	return 1, nil
}

// post author - completed
func (a Authorstore) PostAuthor(ctx context.Context, author entities.Author) (int64, error) {
	res, err := a.db.ExecContext(ctx, "INSERT INTO Author (Id, FirstName,LastName, Dob, PenName) VALUES (?, ? , ?, ?, ?)", author.Id, author.FirstName, author.LastName, author.Dob, author.PenName)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("Error")
	}
	ans, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return ans, nil

	//return http.StatusBadRequest, nil
}

// put author -- completed
func (a Authorstore) PutAuthor(ctx context.Context, author entities.Author, id int) (entities.Author, error) {

	//row := a.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM authors WHERE author_id=?", auth.AuthorID)

	//err := row.Scan(&id)
	//if err != nil {
	//	res, err := a.db.ExecContext(ctx, "INSERT INTO authors(author_id, first_name, last_name, DOB, pen_name) VALUES(?,?,?,?,?)", auth.AuthorID, auth.FirstName, auth.LastName, auth.DOB, auth.PenName)
	//	id, err := res.LastInsertId()
	//
	//	return id, nil
	//} else {
	_, err := a.db.ExecContext(ctx, "UPDATE Author SET author_id=?, first_name=?, last_name=?, DOB=?, pen_name=? WHERE author_id=?",
		author.Id, author.FirstName, author.LastName, author.Dob, author.PenName, id)
	if err != nil {
		return entities.Author{}, err
	}

	return author, nil
	//}

}

// delete author - completed
func (a Authorstore) DeleteAuthor(ctx context.Context, id int) (int64, error) {
	res, err := a.db.ExecContext(ctx, "Delete from Author WHERE Id=?", id)
	if err != nil {
		return 0, nil
	}
	cnt, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return cnt, err

	//return http.StatusNoContent, nil
}
