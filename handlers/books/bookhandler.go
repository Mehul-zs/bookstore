package handlerbook

import (
	"Bookstore/entities"
	"Bookstore/services"
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

//get all books
//func (b Bookhandler) GetAll(rw http.ResponseWriter, req *http.Request) {
//	title := req.URL.Query().Get("Title")
//	getauthor := req.URL.Query().Get("getauthor")
//
//	resp, err := b.serviceBook.GetAll(title, getauthor)
//	if err != nil {
//		_, _ = rw.Write([]byte("Could not post book"))
//		rw.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	body, _ := json.Marshal(resp)
//	_, _ = rw.Write(body)
//
//}

//get book by id
func (b Bookhandler) GetByID(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		log.Fatalln("id is missing")
	}
	//fmt.Println(`id: `, id)

	ID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalln("id not converted into string, error")
	}
	resp, err := b.serviceBook.GetByID(ID)
	if err != nil {
		_, _ = rw.Write([]byte("Could not post book"))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	body, _ := json.Marshal(resp)
	_, _ = rw.Write(body)

}

// post a book
func (b Bookhandler) PostBook(rw http.ResponseWriter, req *http.Request) {
	var book entities.Books
	body, err := io.ReadAll(req.Body)
	if err != nil {
		//fmt.Println("Post handler error")
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write([]byte("invalid Author"))
		return
	}
	json.Unmarshal(body, &book)
	_, err = b.serviceBook.PostBook(book)
	if err != nil {
		fmt.Println(err)
		//fmt.Println("Hiii")
		//_, _ = rw.Write([]byte("Could not post book"))
		rw.WriteHeader(http.StatusInternalServerError)
		_, _ = rw.Write([]byte("Error in put method"))

		return
	}
	rw.WriteHeader(http.StatusCreated)
}

// put a book by id
func (b Bookhandler) PutBook(rw http.ResponseWriter, req *http.Request) {
	var book entities.Books
	body, _ := io.ReadAll(req.Body)
	err := json.Unmarshal(body, &book)
	if err != nil {
		//fmt.Println("Hello handler")
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write([]byte("Invalid book"))
		return
	}
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {

		//fmt.Println("Hello handler")
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write([]byte("Invalid book id"))
		return
	}

	resp, err := b.serviceBook.PutBook(book, id)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		_, _ = rw.Write([]byte("Error in put method"))
		return
	}
	body, _ = json.Marshal(resp)
	rw.WriteHeader(http.StatusAccepted) //202 request accpeted but processing i not completed
	_, _ = rw.Write(body)
}

// delete a book by id
func (b Bookhandler) DeleteBook(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		//fmt.Println("Hello Delete func")
		//_, _ = rw.Write([]byte("invalid parameter id"))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = b.serviceBook.DeleteBook(id)
	if err != nil {
		//_, _ = rw.Write([]byte("could not retrieve book"))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
	return
}
