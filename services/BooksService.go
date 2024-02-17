package services

import (
	application_errors "samuel-gbh-assignment/errors"
	"samuel-gbh-assignment/models"
	"strings"
)

type BooksRepository interface {
	List() ([]models.Book, error)
	FindByID(id int) (models.Book, error)
}

type PagesRepository interface {
	GetPage(bookID, page int) (models.Page, error)
}

type booksService struct {
	booksRepository BooksRepository
	pagesRepository PagesRepository
}

var allowedContentFormats = map[string]bool{
	"text": true,
	"html": true,
}

func (s booksService) List() ([]models.Book, error) {
	books, err := s.booksRepository.List()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s booksService) FindByID(id int) (models.Book, error) {
	book, err := s.booksRepository.FindByID(id)
	if err != nil {
		return models.Book{}, err
	}
	return book, nil
}

func (s booksService) GetPage(bookID, page int, contentFormat string) (string, error) {
	if !allowedContentFormats[strings.ToLower(contentFormat)] {
		return "", application_errors.ErrInvalidContentFormat
	}
	bookPage, err := s.pagesRepository.GetPage(bookID, page)

	if err != nil {
		return "", err
	}
	switch contentFormat {
	case "text":
		return bookPage.ContentTxt, nil
	case "html":
		return bookPage.ContentHTML, nil
	default:
		return "", application_errors.ErrInvalidContentFormat
	}
}

func NewBooksService(booksRepository BooksRepository, pagesRepository PagesRepository) booksService {
	return booksService{
		booksRepository: booksRepository,
		pagesRepository: pagesRepository,
	}
}
