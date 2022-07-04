package handlerAuthor

import (
	"Bookstore/services"
	"net/http"
)

type AuthorHandler struct {
	serviceAuthor services.Author
}

func New(author services.Author) AuthorHandler {
	return AuthorHandler{serviceAuthor: author}
}

func (a AuthorHandler) Handler(rw http.ResponseWriter, req *http.Request) {

}

func (a AuthorHandler) PostAuthor(rw http.ResponseWriter, req *http.Request) {

}

func (a AuthorHandler) PutAuthor(rw http.ResponseWriter, req *http.Request) {

}

func (a AuthorHandler) DeleteAuthor(rw http.ResponseWriter, req *http.Request) {

}
