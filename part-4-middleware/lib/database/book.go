package database

import (
	"restfulapi/ridzqi/config"
	"restfulapi/ridzqi/models"
)

func GetBooks() ([]models.Book, error) {
	var books []models.Book

	if e := config.DB.Find(&books).Error; e != nil {
		return nil, e
	}

	return books, nil
}

func GetBook(id int) (models.Book, error) {
	var book models.Book

	if e := config.DB.First(&book, id).Error; e != nil {
		return book, e
	}

	return book, nil
}

func CreateBook(book models.Book) (models.Book, error) {
	if e := config.DB.Save(&book).Error; e != nil {
		return book, e
	}

	return book, nil
}

func UpdateBook(id int, newBook models.Book) (models.Book, error) {
	var book models.Book

	if e := config.DB.First(&book, id).Error; e != nil {
		return book, e
	}

	book.Title = newBook.Title
	book.Author = newBook.Author

	if e := config.DB.Save(&book).Error; e != nil {
		return book, e
	}

	return book, nil
}

func DeleteBook(id int) (models.Book, error) {
	var book models.Book

	if e := config.DB.First(&book, id).Error; e != nil {
		return book, e
	}

	if e := config.DB.Delete(&book).Error; e != nil {
		return book, e
	}

	return book, nil
}
