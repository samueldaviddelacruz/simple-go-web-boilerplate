package tests_mocks

import (
	application_errors "samuel-gbh-assignment/errors"
	"samuel-gbh-assignment/models"
)

type BooksRepositoryMock struct{}

func (b BooksRepositoryMock) List() ([]models.Book, error) {
	return []models.Book{{Id: 1, Title: "test", Author: "test", Numpages: 1}}, nil
}

func (b BooksRepositoryMock) FindByID(id int) (models.Book, error) {
	if id != 1 {
		return models.Book{}, application_errors.ErrBookNotFound
	}
	return models.Book{Id: 1, Title: "test", Author: "test", Numpages: 1}, nil
}

type PagesRepositoryMock struct{}

func (p PagesRepositoryMock) GetPage(bookID, page int) (models.Page, error) {
	if bookID != 1 || page != 1 {
		return models.Page{}, application_errors.ErrPageNotFound
	}
	return models.Page{Bookid: 1, Number: 1, ContentTxt: "test", ContentHTML: "test"}, nil
}

type BooksServiceMock struct{}

func (b BooksServiceMock) List() ([]models.Book, error) {
	return []models.Book{{Id: 1, Title: "test", Author: "test", Numpages: 1}}, nil
}

func (b BooksServiceMock) FindByID(id int) (models.Book, error) {
	if id != 1 {
		return models.Book{}, application_errors.ErrBookNotFound
	}
	return models.Book{Id: 1, Title: "test", Author: "test", Numpages: 1}, nil
}
func (b BooksServiceMock) GetPage(bookID, page int, contentFormat string) (string, error) {
	if bookID != 1 || page != 1 || contentFormat != "text" {
		return "", application_errors.ErrPageNotFound
	}
	return "test", nil
}
