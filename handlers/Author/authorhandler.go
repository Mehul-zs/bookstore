package handlerauthor

import (
	"Bookstore/entities"
	"Bookstore/services"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

type AuthorHandler struct {
	serviceAuthor services.Author
}

func New(author services.Author) AuthorHandler {
	return AuthorHandler{serviceAuthor: author}
}

// post author  - completed
func (a AuthorHandler) PostAuthor(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	var author entities.Author
	body, err := io.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &author)
	if err != nil {
		fmt.Println("hello to post handler")
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write([]byte("Invalid Author"))
		return
	}

	res, err := a.serviceAuthor.PostAuthor(ctx, author)
	if err != nil {
		rw.Write([]byte("Could not post author"))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	body, err = json.Marshal(res)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Write(body)

}

// put author by id  - completed
func (a AuthorHandler) PutAuthor(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	var author entities.Author
	body, err := io.ReadAll(req.Body)
	if err != nil {
		//fmt.Println("Hello handler")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Invalid Author"))
		return
	}
	err = json.Unmarshal(body, &author)
	if err != nil {

	}

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := a.serviceAuthor.PutAuthor(ctx, author, id)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error in put method"))
		return
	}
	body, _ = json.Marshal(res)
	rw.Write(body)

}

// delete author by id  - commpleted
func (a AuthorHandler) DeleteAuthor(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write([]byte("Invalid Author"))
		return
	}

	_, err = a.serviceAuthor.DeleteAuthor(ctx, id)
	if err != nil {
		rw.Write([]byte("Error in Delete method"))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
