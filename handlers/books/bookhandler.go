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

//get all books  - completed
func (b Bookhandler) GetAllBooks(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	title := req.URL.Query().Get("Title")
	getauthor := req.URL.Query().Get("getauthor")

	res, err := b.serviceBook.GetAllBooks(ctx, title, getauthor)
	if err != nil {
		rw.Write([]byte("Could not post book"))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	body, err := json.Marshal(res)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write(body)

}

//get book by id  - completed
func (b Bookhandler) GetBookByID(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatalln("id not converted into string, error")
	}
	res, err := b.serviceBook.GetBookByID(ctx, id)
	if err != nil {
		_, _ = rw.Write([]byte("Could not post book"))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	body, err := json.Marshal(res)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(body)

}

// post a book  - completed
func (b Bookhandler) PostBook(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	var book entities.Book
	body, err := io.ReadAll(req.Body)
	if err != nil {
		//fmt.Println("Post handler error")
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write([]byte("invalid Author"))
		return
	}
	err = json.Unmarshal(body, &book)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := b.serviceBook.PostBook(ctx, book)
	if err != nil {
		fmt.Println(err)
		//fmt.Println("Hiii")
		//_, _ = rw.Write([]byte("Could not post book"))
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error in put method"))

		return
	}

	//rw.WriteHeader(http.StatusCreated)
	res, err := json.Marshal(resp)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}
	rw.Write(res)
}

// put a book by id   - completed
func (b Bookhandler) PutBook(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	var book entities.Book
	body, err := io.ReadAll(req.Body)

	if err != nil {
		log.Fatalln("Error in data")
	}

	err = json.Unmarshal(body, &book)
	if err != nil {
		//fmt.Println("Hello handler")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//fmt.Println("Hello handler")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Invalid book id"))
		return
	}

	resp, err := b.serviceBook.PutBook(ctx, book, id)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error in put method"))
		return
	}
	body, err = json.Marshal(resp)
	if err != nil {
		log.Fatalln("Marshalling error in put handler")
	}
	//rw.WriteHeader(http.StatusAccepted) //202 request accpeted but processing i not completed
	rw.Write(body)
}

// delete a book by id  - completed
func (b Bookhandler) DeleteBook(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		//fmt.Println("Hello Delete func")
		//_, _ = rw.Write([]byte("invalid parameter id"))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = b.serviceBook.DeleteBook(ctx, id)
	if err != nil {
		//_, _ = rw.Write([]byte("could not retrieve book"))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
	return
}
