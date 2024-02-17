package controllers

import (
	"errors"
	"fmt"
	"net/http"
	application_errors "simple-go-web-boilerplate/errors"
	"simple-go-web-boilerplate/models"
	"strconv"
	"strings"

	"encoding/json"

	"github.com/go-chi/chi/v5"
)

var ContentTypeMapping = map[string]string{
	"text": "text/plain",
	"html": "text/html",
}

type BooksService interface {
	List() ([]models.Book, error)
	FindByID(id int) (models.Book, error)
	GetPage(bookID, page int, contentFormat string) (string, error)
}
type booksController struct {
	booksService BooksService
}

func (h booksController) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books, err := h.booksService.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error getting the books"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	booksJson, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error parsing the books"}`))
		return
	}
	w.Write(booksJson)
}
func (h booksController) FindByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	w.Header().Set("Content-Type", "application/json")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Error parsing the id"}`))
	}

	book, err := h.booksService.FindByID(bookID)
	if err != nil {
		if errors.Is(err, application_errors.ErrBookNotFound) {
			fmt.Println("error", err)
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	bookJson, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error parsing the book"}`))
		return
	}
	w.Write(bookJson)
}
func (h booksController) GetPage(w http.ResponseWriter, r *http.Request) {
	bookID := chi.URLParam(r, "bookid")
	page := chi.URLParam(r, "page")
	contentFormat := chi.URLParam(r, "contentFormat")
	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Error parsing the bookid"}`))
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Error parsing the page"}`))
	}

	pageContent, err := h.booksService.GetPage(bookIDInt, pageInt, contentFormat)

	if err != nil {
		switch err {
		case application_errors.ErrPageNotFound:
			w.WriteHeader(http.StatusNotFound)
		case application_errors.ErrInvalidContentFormat:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", ContentTypeMapping[strings.ToLower(contentFormat)])
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(pageContent))
}
func NewBooksController(booksService BooksService) booksController {
	return booksController{
		booksService: booksService,
	}
}
