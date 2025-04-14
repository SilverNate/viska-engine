package library

import (
	"gorm.io/gorm"
)

type LibraryRepository struct {
	db *gorm.DB
}

func NewLibraryRepository(db *gorm.DB) *LibraryRepository {
	return &LibraryRepository{db}
}

func (r *LibraryRepository) CreateAuthor(author *Author) error {
	err := r.db.Create(author).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *LibraryRepository) FindAllAuthor() ([]Author, error) {
	var authors []Author
	err := r.db.Preload("Books").Find(&authors).Error
	if err != nil {
		return authors, err
	}
	return authors, err
}

func (r *LibraryRepository) FindAuthorByID(id uint) (*Author, error) {
	var author Author
	err := r.db.Preload("Books").First(&author, id).Error
	if err != nil {
		return &author, err
	}
	return &author, err
}

func (r *LibraryRepository) UpdateAuthor(author *Author) error {
	err := r.db.Save(author).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *LibraryRepository) DeleteAuthor(id uint) error {
	err := r.db.Delete(&Author{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *LibraryRepository) CreateBook(books *Book) error {
	err := r.db.Create(books).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *LibraryRepository) FindAllBook() ([]Book, error) {
	var books []Book
	err := r.db.Preload("Author").Preload("Publisher").Find(&books).Error
	if err != nil {
		return books, err
	}
	return books, err
}

func (r *LibraryRepository) FindBookByID(id uint) (*Book, error) {
	var book Book
	err := r.db.Preload("Author").Preload("Publisher").First(&book, id).Error
	if err != nil {
		return &book, err
	}
	return &book, err
}

func (r *LibraryRepository) UpdateBook(books *Book) error {
	err := r.db.Save(books).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *LibraryRepository) DeleteBook(id uint) error {
	err := r.db.Delete(&Book{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *LibraryRepository) CreatePublisher(publishers *Publisher) error {
	err := r.db.Create(publishers).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *LibraryRepository) FindAllPublisher() ([]Publisher, error) {
	var publishers []Publisher
	err := r.db.Preload("Books").Find(&publishers).Error
	if err != nil {
		return publishers, err
	}
	return publishers, err
}

func (r *LibraryRepository) FindPublisherByID(id uint) (*Publisher, error) {
	var publishers Publisher
	err := r.db.Preload("Books").First(&publishers, id).Error
	if err != nil {
		return &publishers, err
	}
	return &publishers, err
}

func (r *LibraryRepository) UpdatePublisher(publishers *Publisher) error {
	err := r.db.Save(publishers).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *LibraryRepository) DeletePublisher(id uint) error {
	err := r.db.Delete(&Publisher{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
