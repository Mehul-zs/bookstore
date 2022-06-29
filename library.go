package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Books struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Author        Author `json:"author"`
	Publication   string `json:"publication"`
	PublishedDate string `json:"published_date"`
}

type Author struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Dob       string `json:"dob"`
	PenName   string `json:"pen_name"`
}

func DbConnect() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "HelloMehul1@"
	dbName := "bookstore"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)
	//fmt.Println("hello Mehul")
	if err != nil {
		panic(err.Error())
	}
	//fmt.Println("hello Charlie")
	return db
}

func BookExists(book Books, Id int64) bool {

	db := DbConnect()
	res, err := db.Query("SELECT Id FROM Book WHERE Title = ? AND Publication =? AND  PublsihedDate= ? AND AuthorId =?", book.Title, book.Publication, book.PublishedDate, Id)
	if err != nil || !res.Next() {
		return false
	}

	return true
}

func AuthorIDExists(Id int) bool {
	db := DbConnect()
	res, err := db.Query("Select Id from Author WHERE Id=?", Id)
	if !res.Next() || err != nil {
		return false
	}
	return true
}

func GetBook(rw http.ResponseWriter, req *http.Request) {
	//json.NewEncoder(rw).Encode([]Books{{1, "Mehul", nil, "Penguin", "10/08/2000"}, {2, "Mehul", nil, "Penguin", "10/08/2000"}}) // for manual testing
	db := DbConnect()
	title := req.URL.Query().Get("Title")
	var rows *sql.Rows
	var err error
	if title == "" {
		rows, err = db.Query("select * from Books;")
	} else {
		rows, err = db.Query("select * from Books where title=?;", title)
	}
	if err != nil {
		log.Print(err)
	}
	var books []Books
	for rows.Next() {
		book := Books{}
		err = rows.Scan(&book.Id, &book.Title, &book.Publication, &book.PublishedDate)
		if err != nil {
			log.Print(err)
		}
		getAuthor := req.URL.Query().Get("GetAuthor")
		if getAuthor == "true" {
			row := db.QueryRow("SELECT * from Author WHERE Id=? ", book.Author.Id)
			row.Scan(&book.Author.Id, &book.Author.FirstName, &book.Author.LastName, &book.Author.Dob, &book.Author.PenName)
		}
		books = append(books, book)
	}
	json.NewEncoder(rw).Encode(books)

}

func GetBookById(rw http.ResponseWriter, req *http.Request) {
	db := DbConnect()
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing")
	}
	fmt.Println(`id: `, id)

	var rows *sql.Rows
	var err error
	if id == "" {
		rows, err = db.Query("select * from Books;", id)
	} else {
		rows, err = db.Query("select * from Books where id=?;", id)

	}
	if err != nil {
		log.Print(err)
	}
	var books []Books
	for rows.Next() {
		book := Books{}
		err = rows.Scan(&book.Id, &book.Title, &book.Publication, &book.PublishedDate)
		if err != nil {
			log.Print(err)
		}

		books = append(books, book)
	}
	json.NewEncoder(rw).Encode(books)

}

func PostByAuthor(rw http.ResponseWriter, req *http.Request) {
	//json.NewEncoder(rw).Encode([]Books{{1, "Mehul", nil, "Penguin", "10/08/2000"}, {2, "Mehul", nil, "Penguin", "10/08/2000"}}) // for manual testing
	fmt.Println("Hello")
	db := DbConnect()
	var author *Author
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
		return

	}

	err = json.Unmarshal(body, &author)
	if err != nil {
		log.Print(err)
		return
	}

	fmt.Println(author.Id)
	if author.FirstName == "" || author.LastName == "" {
		rw.WriteHeader(http.StatusBadRequest)

		return
	}
	//fmt.Println(db.Query("SELECT * FROM Author"))
	res, err := db.Query("SELECT Id FROM Author WHERE FirstName=? AND LastName=? AND  Dob=? AND PenName=?", author.FirstName, author.LastName, author.Dob, author.PenName)
	fmt.Println(res.Next())
	if !res.Next() || err != nil {
		_, err = db.Exec("INSERT INTO Author (Id, FirstName,LastName, Dob, PenName) VALUES (?, ? , ?, ?, ?)", author.Id, author.FirstName, author.LastName, author.Dob, author.PenName)
		if err != nil {
			return
		}
		rw.WriteHeader(http.StatusCreated)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func PostByBook(rw http.ResponseWriter, req *http.Request) {
	db := DbConnect()
	var Book Books
	body, err := io.ReadAll(req.Body)
	//fmt.Println(body)
	//fmt.Println(Book)

	if err != nil {
		log.Print(err)
	}
	fmt.Println(db)
	err = json.Unmarshal(body, &Book)
	if err != nil {
		log.Print(err)
		return
	}

	fmt.Println(Book)

	if Book.Title == "" || Book.Author.FirstName == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if Book.Publication != "Scholastic" && Book.Publication != "Arihant" && Book.Publication != "Penguin" {
		log.Print("Invalid Publication")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	publishedDate := strings.Split(Book.PublishedDate, "/")
	if len(publishedDate) < 3 {
		return
	}
	yr, _ := strconv.Atoi(publishedDate[2])

	if yr >= 2022 || yr < 1880 {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	var authId int64
	result, err := db.Query("SELECT Id FROM Author WHERE FirstName = ? AND LastName =? AND  Dob = ? AND PenName =?", Book.Author.FirstName, Book.Author.LastName, Book.Author.Dob, Book.Author.PenName)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(result)
	if result.Next() {
		err = result.Scan(&authId)
		if err != nil {
			log.Print("author is not present")
		}
	} else {
		res, err := db.Exec("INSERT INTO Author (Id, FirstName,LastName,Dob,PenName) VALUES (?, ?, ?, ?, ?)", Book.Author.Id, Book.Author.FirstName, Book.Author.LastName, Book.Author.Dob, Book.Author.PenName)
		if err != nil {
			return
		}
		authId, err = res.LastInsertId()
	}
	fmt.Println("last, Hello Mehul")

	if !BookExists(Book, authId) {
		_, err = db.Exec("INSERT INTO Books (Id, Title, Publication,PublishedDate, AuthorId) VALUES (? ,?, ?, ?, ?)", Book.Id, Book.Title, Book.Publication, Book.PublishedDate, authId)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		rw.WriteHeader(http.StatusCreated)
	}
	//rw.WriteHeader(http.StatusCreated)

}

func PutBook(rw http.ResponseWriter, req *http.Request) {
	db := DbConnect()
	var book Books
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &book)
	if err != nil {
		return
	}
	if book.Title == "" {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Println("test case failed at books title: %s", book.Title)
		return
	}
	if !(book.Publication == "Penguin" || book.Publication == "Arihant" || book.Publication == "Scholastic") {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Println("test case failed at books publication: %s", book.Publication)
		return
	}

	publicationDate := strings.Split(book.PublishedDate, "/")
	if len(publicationDate) < 3 {
		return
	}
	yr, _ := strconv.Atoi(publicationDate[2])
	if yr > time.Now().Year() || yr < 1880 {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Println("Test case failed at published data")
		return
	}

	params := mux.Vars(req)
	ID, err := strconv.Atoi(params["id"])
	if ID <= 0 {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := db.Query("SELECT Id FROM Author WHERE Id = ?", ID)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if !result.Next() {
		log.Print("author not present", book.Author.Id)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err = db.Query("SELECT * FROM Books WHERE Id = ?", ID)
	if err != nil {
		log.Print(err)
	}
	if !result.Next() {
		log.Print("Book not present")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err = db.Query("UPDATE Books SET Id=?, Title = ? ,Publication = ? ,PublishedDate=?  WHERE id =?", ID, book.Title, book.Publication, book.PublishedDate, book.Id)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	} else {
		rw.WriteHeader(http.StatusCreated)
	}

}

func PutAuthor(rw http.ResponseWriter, req *http.Request) {

	db := DbConnect()
	var author Author
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
		return
	}
	err = json.Unmarshal(body, &author)
	if err != nil {
		log.Print(err)
		return
	}
	if author.FirstName == "" || author.LastName == "" || author.PenName == "" || author.Dob == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	params := mux.Vars(req)
	ID, err := strconv.Atoi(params["id"])
	if ID <= 0 {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := db.Query("SELECT Id FROM Author WHERE Id = ?", ID)
	if err != nil {
		log.Print(err)
	}
	if !res.Next() {
		log.Print("id not present")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var id int
	err = res.Scan(&id)
	if err != nil {
		log.Print(err)
		return
	}

	_, err = db.Exec("UPDATE Author SET FirstName = ?, LastName = ?, Dob = ? , PenName = ?, Id=?  WHERE Id =?", author.FirstName, author.LastName, author.Dob, author.PenName, ID, ID)
	if err != nil {
		log.Print(err)
		return
	}
	rw.WriteHeader(http.StatusCreated)

}

func DeleteBook(rw http.ResponseWriter, req *http.Request) {

	db := DbConnect()
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	//ok status ?
	if err != nil {
		fmt.Println("Invalid Id")
	}

	if id <= 0 {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	//
	_, err = db.Exec("Delete from Books WHERE Id=?", id)
	if err != nil {
		log.Print("book not exist")
	}
	rw.WriteHeader(http.StatusNoContent)
}

func DeleteAuthor(rw http.ResponseWriter, req *http.Request) {
	db := DbConnect()
	params := mux.Vars(req)

	Id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Id is missing/Invalid Id")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if Id <= 0 {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if !AuthorIDExists(Id) {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Println("Author does not exists")
		return
	}

	check, err := db.Query("Select Id from Books WHERE AuthorId=?", Id)

	for check.Next() {
		var id int
		err := check.Scan(&id)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = db.Exec("Delete from Books WHERE Id=?", id)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	_, err = db.Exec("Delete from Author WHERE Id=?", Id)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

func main() {

	r := mux.NewRouter()
	db := DbConnect()
	fmt.Println(db)
	r.HandleFunc("/books", GetBook).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", GetBookById).Methods(http.MethodGet)

	r.HandleFunc("/author", PostByAuthor).Methods(http.MethodPost)
	r.HandleFunc("/book", PostByBook).Methods(http.MethodPost)

	r.HandleFunc("/author/{id}", PutAuthor).Methods(http.MethodPut)
	r.HandleFunc("/books/{id}", PutBook).Methods(http.MethodPut)
	//
	r.HandleFunc("/deleteBook/{id}", DeleteBook).Methods(http.MethodDelete)
	r.HandleFunc("/deleteAuthor/{id}", DeleteAuthor).Methods(http.MethodDelete)
	fmt.Println("Hey Mehul!")
	//r.HandleFunc("/PutBook")

	Server := http.Server{
		Addr:    ":5000",
		Handler: r,
	}

	fmt.Println("Server started at 5000")
	log.Fatal(Server.ListenAndServe())

}
