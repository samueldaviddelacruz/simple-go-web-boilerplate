package repositories

import (
	"database/sql"
	"errors"
	application_errors "simple-go-web-boilerplate/errors"
	"simple-go-web-boilerplate/models"

	"github.com/jmoiron/sqlx"
)

type pagesRepository struct {
	db *sqlx.DB
}

func (r pagesRepository) GetPage(bookID, pageNumber int) (models.Page, error) {
	var p models.Page
	err := r.db.Get(&p, "SELECT * FROM pages WHERE bookId = ? AND number = ?", bookID, pageNumber)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Page{}, application_errors.ErrPageNotFound
		}
		return models.Page{}, application_errors.ErrInternalServerError
	}
	return p, nil
}

func NewPagesRepository(db *sqlx.DB) pagesRepository {
	return pagesRepository{db}
}
