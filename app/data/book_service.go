package data

import "github.com/jinzhu/gorm"

type BookService interface {
	Book(id uint) *Book
	PostBook(title, author string) (uint, error)
	PutBook(id uint, title, author string) error
	DeleteBook(id uint) error
}

func (database *Database) Book(id uint) *Book {
	var book Book
	if database.First(&book, id).RecordNotFound() {
		return nil
	}

	return &book
}

func (database *Database) PostBook(title, author string) (uint, error) {
	book := Book{
		Title:  title,
		Author: author,
	}
	err := database.Create(&book).Error

	return book.ID, err
}

func (database *Database) PutBook(id uint, title, author string) error {
	var book Book
	if database.First(&book, id).RecordNotFound() {
		return ErrResourceNotFound
	}

	book.Title = title
	book.Author = author

	return database.Save(&book).Error
}

func (database *Database) DeleteBook(id uint) error {
	return database.Delete(&Book{Model: gorm.Model{ID: id}}).Error
}
