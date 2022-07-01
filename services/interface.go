package services

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
