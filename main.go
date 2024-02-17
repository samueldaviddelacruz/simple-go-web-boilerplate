package main

import (
	"fmt"
	"log"
	"net/http"
	controllers "samuel-gbh-assignment/controllers/books"
	"samuel-gbh-assignment/repositories"
	"samuel-gbh-assignment/services"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := sqlx.Open("sqlite", "./data/books.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	booksRepository := repositories.NewBooksRepository(db)
	pagesRepository := repositories.NewPagesRepository(db)
	booksService := services.NewBooksService(booksRepository, pagesRepository)
	booksController := controllers.NewBooksController(booksService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/books", booksController.List)
	r.Get("/books/{id}", booksController.FindByID)
	r.Get("/books/{bookid}/pages/{page}/{contentFormat}", booksController.GetPage)

	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", r)
}
