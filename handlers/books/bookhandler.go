package handlerBook

import (
	"Bookstore/services"
	"net/http"
)

type Bookhandler struct {
	serviceBook services.Book
}

func New(s interface{}) Bookhandler {
	return Bookhandler{serviceBook: s}
}

func (b Bookhandler) Handler(rw http.ResponseWriter, req *http.Request) {

}

func (b Bookhandler) GetAll(rw http.ResponseWriter, req *http.Request) {

}

func (b Bookhandler) GetByID(rw http.ResponseWriter, req *http.Request) {

}

func (b Bookhandler) PostBook(rw http.ResponseWriter, req *http.Request) {

}

func (b Bookhandler) PutBook(rw http.ResponseWriter, req *http.Request) {

}

func (b Bookhandler) DeleteBook(rw http.ResponseWriter, req *http.Request) {

}
