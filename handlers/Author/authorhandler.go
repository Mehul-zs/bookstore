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

// post author
func (a AuthorHandler) PostAuthor(rw http.ResponseWriter, req *http.Request) {
	var author entities.Author
	body, _ := io.ReadAll(req.Body)
	err := json.Unmarshal(body, &author)

	if err != nil {
		fmt.Println("hello to post handler")
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write([]byte("Invalid Author"))
		return
	}

	resp, err := a.serviceAuthor.PostAuthor(author)
	if err != nil {
		_, _ = rw.Write([]byte("Could not post author"))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	body, _ = json.Marshal(resp)

	rw.WriteHeader(http.StatusCreated)
	_, _ = rw.Write(body)

}

// put author by id
func (a AuthorHandler) PutAuthor(rw http.ResponseWriter, req *http.Request) {
	var author entities.Author
	body, _ := io.ReadAll(req.Body)
	err := json.Unmarshal(body, &author)
	if err != nil {
		fmt.Println("Hello handler")
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write([]byte("Invalid Author"))
		return
	}

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := a.serviceAuthor.PutAuthor(author, id)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		_, _ = rw.Write([]byte("Error in put method"))
		return
	}
	body, _ = json.Marshal(resp)
	_, _ = rw.Write(body)

}

// delete author by id
func (a AuthorHandler) DeleteAuthor(rw http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write([]byte("Invalid Author"))
		return
	}

	_, err = a.serviceAuthor.DeleteAuthor(id)
	if err != nil {
		_, _ = rw.Write([]byte("Error in Delete method"))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
