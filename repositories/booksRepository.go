package repositories

import (
	"database/sql"
	"errors"
	application_errors "samuel-gbh-assignment/errors"
	"samuel-gbh-assignment/models"

	"github.com/jmoiron/sqlx"
)

type booksRepository struct {
	db *sqlx.DB
}

func (r booksRepository) List() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Select(&books, "SELECT * FROM books")
	if err != nil {
		return []models.Book{}, application_errors.ErrInternalServerError
	}
	return books, nil
}

func (r booksRepository) FindByID(id int) (models.Book, error) {
	var book models.Book
	err := r.db.Get(&book, "SELECT * FROM books WHERE id = ?", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Book{}, application_errors.ErrBookNotFound
		}
		return models.Book{}, application_errors.ErrInternalServerError
	}
	return book, nil
}

func NewBooksRepository(db *sqlx.DB) booksRepository {
	return booksRepository{db}
}
