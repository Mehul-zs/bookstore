package datastoreauthor

import (
	"Bookstore/entities"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

type Authorstore struct {
	db *sql.DB
}

func New(db *sql.DB) Authorstore {
	return Authorstore{db: db}
}

func (a Authorstore) PostAuthor(author entities.Author) (int64, error) {
	res, err := a.db.Query("SELECT Id FROM Author WHERE FirstName=? AND LastName=? AND  Dob=? AND PenName=? AND Id =?", author.FirstName, author.LastName, author.Dob, author.PenName, author.Id)
	if err != nil {
		return http.StatusBadRequest, errors.New("Author Alreadyexists")

	}
	//fmt.Println(res.Next())
	if !res.Next() || err != nil {
		_, err := a.db.Exec("INSERT INTO Author (Id, FirstName,LastName, Dob, PenName) VALUES (?, ? , ?, ?, ?)", author.Id, author.FirstName, author.LastName, author.Dob, author.PenName)
		if err != nil {
			fmt.Println("Hello to db")
			return http.StatusBadRequest, nil
		}
		//_, err = ans.LastInsertId()
		return http.StatusCreated, nil

	}
	return http.StatusBadRequest, nil
}

func (a Authorstore) PutAuthor(author entities.Author, id int) (entities.Author, error) {
	res, err := a.db.Query("SELECT Id FROM Author WHERE FirstName=? AND LastName=? AND  Dob=? AND PenName=?", author.FirstName, author.LastName, author.Dob, author.PenName)
	if !res.Next() || err != nil {
		_, err = a.db.Exec("UPDATE Author SET FirstName = ?, LastName = ?, Dob = ? , PenName = ?, Id=?  WHERE Id =?", author.FirstName, author.LastName, author.Dob, author.PenName, author.Id, author.Id)

		if err != nil {
			//fmt.Println("Put author error")
			return entities.Author{}, nil
		}
		//fmt.Println("Hello mehul put author")
		return author, nil
	}
	//fmt.Println("Hello")
	return author, nil
}

func (a Authorstore) DeleteAuthor(id int) (int64, error) {

	check, err := a.db.Query("Select Id from Books WHERE AuthorId=?", id)
	if err != nil {
		return http.StatusBadRequest, nil
	}

	for check.Next() {
		var ID int
		err = check.Scan(&ID)
		if err != nil {
			return http.StatusBadRequest, nil
		}
		_, err = a.db.Exec("Delete from Books WHERE Id=?", ID)
		if err != nil {
			return http.StatusBadRequest, nil
		}
	}

	_, err = a.db.Exec("Delete from Author WHERE Id=?", id)
	if err != nil {
		return http.StatusBadRequest, nil
	}

	return http.StatusNoContent, nil
}
