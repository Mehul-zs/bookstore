package handlerbook

import (
	"Bookstore/entities"
	"Bookstore/services"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Bookhandler struct {
	serviceBook services.Book
}

func New(s services.Book) Bookhandler {
	return Bookhandler{serviceBook: s}
}

//get all books  - completed
func (b Bookhandler) GetAllBooks(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	title := req.URL.Query().Get("title")
	getauthor := req.URL.Query().Get("getauthor")

	res, err := b.serviceBook.GetAllBooks(ctx, title, getauthor)
	if err != nil {
		fmt.Println(err)
		rw.Write([]byte("Could not post book"))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	body, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write(body)

}

//get book by id  - completed running properly
func (b Bookhandler) GetBookByID(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("id not converted into string, error")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := b.serviceBook.GetBookByID(context.Background(), id)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	body, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	//rw.WriteHeader(http.StatusOK)
	rw.Write(body)
}

func (b Bookhandler) PostBook(rw http.ResponseWriter, req *http.Request) {
	var book entities.Book
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Hello")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("invalid Author"))
		return
	}
	err = json.Unmarshal(body, &book)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := b.serviceBook.PostBook(context.Background(), book)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error in put method"))
		return
	}

	_, err = json.Marshal(resp)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}
	rw.WriteHeader(http.StatusCreated)
	//rw.Write(res)
}

func (b Bookhandler) PutBook(rw http.ResponseWriter, req *http.Request) {
	var book entities.Book
	body, err := io.ReadAll(req.Body)

	if err != nil {
		log.Println("Error in data")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &book)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Invalid book id"))
		return
	}

	resp, err := b.serviceBook.PutBook(context.Background(), book, id)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error in put method"))
		return
	}
	body, err = json.Marshal(resp)
	if err != nil {
		log.Fatalln("Marshalling error in put handler")
	}
	rw.WriteHeader(http.StatusAccepted) //202 request accpeted but processing i not completed
	rw.Write(body)
}

func (b Bookhandler) DeleteBook(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := b.serviceBook.DeleteBook(context.Background(), id)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if res != 0 {
		rw.WriteHeader(http.StatusNoContent)
		//rw.Write([]byte("Book deleted"))
		log.Println("book deleted")
		return
	}
	rw.WriteHeader(http.StatusBadRequest)
	return
}
