package datastore

type Book interface {
	GetAll()
	GetByID()
	PostBOok()
	DeleteBook()
	PutBook()
}

type Author interface {
	PostAuthor()
	DeleteAuthor()
	PutAuthor()
}
