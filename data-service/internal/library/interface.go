package library

//go:generate mockgen -destination=mocks/mock.go -package=mocks -source=interface.go
type ILibraryService interface {
	CreateAuthor(author *Author) error
	GetAllAuthor() ([]Author, error)
	GetAuthorByID(id uint) (*Author, error)
	UpdateAuthor(author *Author) error
	DeleteAuthor(id uint) error

	CreatePublisher(author *Publisher) error
	GetAllPublisher() ([]Publisher, error)
	GetPublisherByID(id uint) (*Publisher, error)
	UpdatePublisher(author *Publisher) error
	DeletePublisher(id uint) error

	CreateBook(author *Book) error
	GetAllBook() ([]Book, error)
	GetBookByID(id uint) (*Book, error)
	UpdateBook(author *Book) error
	DeleteBook(id uint) error
}

type ILibraryRepo interface {
	CreateAuthor(author *Author) error
	FindAllAuthor() ([]Author, error)
	FindAuthorByID(id uint) (*Author, error)
	UpdateAuthor(author *Author) error
	DeleteAuthor(id uint) error

	CreatePublisher(publisher *Publisher) error
	FindAllPublisher() ([]Publisher, error)
	FindPublisherByID(id uint) (*Publisher, error)
	UpdatePublisher(publisher *Publisher) error
	DeletePublisher(id uint) error

	CreateBook(publisher *Book) error
	FindAllBook() ([]Book, error)
	FindBookByID(id uint) (*Book, error)
	UpdateBook(publisher *Book) error
	DeleteBook(id uint) error
}
