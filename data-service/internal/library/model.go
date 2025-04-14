package library

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Books []Book `gorm:"foreignKey:AuthorID"`
}

type Book struct {
	gorm.Model
	Title       string    `gorm:"not null"`
	AuthorID    uint      `json:"authorId"`
	Author      Author    `gorm:"foreignKey:AuthorID"`
	PublisherID uint      `json:"publisherId"`
	Publisher   Publisher `gorm:"foreignKey:PublisherID"`
}

type Publisher struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Books []Book `gorm:"foreignKey:PublisherID"`
}
