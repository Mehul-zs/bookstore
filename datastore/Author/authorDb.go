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

func (a Authorstore) GetAllAuthor(ctx context.Context) ([]entities.Author, error) {
	rows, err := a.db.QueryContext(ctx, "SELECT * FROM Author")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var authors []entities.Author
	for rows.Next() {
		var author entities.Author
		err = rows.Scan(author.Id, author.FirstName, author.LastName, author.Dob, author.PenName)
		if err != nil {
			return []entities.Author{}, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}

func (a Authorstore) CheckAuthor(ctx context.Context, author entities.Author) (int, error) {
	res, err := a.db.QueryContext(ctx, "SELECT Id FROM Author WHERE FirstName=? AND LastName=? AND  Dob=? AND PenName=? AND Id =?",
		author.FirstName, author.LastName, author.Dob, author.PenName, author.Id)
	if err != nil {
		return 0, errors.New("author alreadyexists")
	}
	if res.Next() {
		fmt.Println("Hello data")
		return 1, nil
	}
	return 0, nil
}

func (a Authorstore) PostAuthor(ctx context.Context, author entities.Author) (int64, error) {
	res, err := a.db.ExecContext(ctx, "INSERT INTO Author (Id, FirstName,LastName, Dob, PenName) VALUES (?, ? , ?, ?, ?)",
		author.Id, author.FirstName, author.LastName, author.Dob, author.PenName)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("error")
	}
	ans, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return ans, nil
}

func (a Authorstore) PutAuthor(ctx context.Context, author entities.Author, id int) (entities.Author, error) {
	_, err := a.db.ExecContext(ctx, "UPDATE Author SET author_id=?, first_name=?, last_name=?, DOB=?, pen_name=? WHERE author_id=?",
		author.Id, author.FirstName, author.LastName, author.Dob, author.PenName, id)
	if err != nil {
		return entities.Author{}, err
	}
	return author, nil
}

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
}
