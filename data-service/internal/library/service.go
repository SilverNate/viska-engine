package library

type LibraryService struct {
	repo ILibraryRepo
}

func NewLibraryService(repo ILibraryRepo) *LibraryService {
	return &LibraryService{repo: repo}
}

func (s *LibraryService) CreateAuthor(author *Author) error {
	return s.repo.CreateAuthor(author)
}

func (s *LibraryService) GetAllAuthor() ([]Author, error) {
	return s.repo.FindAllAuthor()
}

func (s *LibraryService) GetAuthorByID(id uint) (*Author, error) {
	return s.repo.FindAuthorByID(id)
}

func (s *LibraryService) UpdateAuthor(author *Author) error {
	return s.repo.UpdateAuthor(author)
}

func (s *LibraryService) DeleteAuthor(id uint) error {
	return s.repo.DeleteAuthor(id)
}

func (s *LibraryService) CreatePublisher(publisher *Publisher) error {
	return s.repo.CreatePublisher(publisher)
}

func (s *LibraryService) GetAllPublisher() ([]Publisher, error) {
	return s.repo.FindAllPublisher()
}

func (s *LibraryService) GetPublisherByID(id uint) (*Publisher, error) {
	return s.repo.FindPublisherByID(id)
}

func (s *LibraryService) UpdatePublisher(publisher *Publisher) error {
	return s.repo.UpdatePublisher(publisher)
}

func (s *LibraryService) DeletePublisher(id uint) error {
	return s.repo.DeletePublisher(id)
}

func (s *LibraryService) CreateBook(book *Book) error {
	return s.repo.CreateBook(book)
}

func (s *LibraryService) GetAllBook() ([]Book, error) {
	return s.repo.FindAllBook()
}

func (s *LibraryService) GetBookByID(id uint) (*Book, error) {
	return s.repo.FindBookByID(id)
}

func (s *LibraryService) UpdateBook(book *Book) error {
	return s.repo.UpdateBook(book)
}

func (s *LibraryService) DeleteBook(id uint) error {
	return s.repo.DeleteBook(id)
}
