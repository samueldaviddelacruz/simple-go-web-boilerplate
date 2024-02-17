package controllers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	controllers "simple-go-web-boilerplate/controllers/books"
	tests_mocks "simple-go-web-boilerplate/tests"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestBooksControllerFindById(t *testing.T) {
	booksController := controllers.NewBooksController(tests_mocks.BooksServiceMock{})
	testCases := []struct {
		name           string
		bookId         string
		expectedStatus int
	}{
		{name: "Test FindByID", bookId: "1", expectedStatus: http.StatusOK},
		{name: "Test FindByID with invalid id", bookId: "asdsa", expectedStatus: http.StatusBadRequest},
		{name: "Test FindByID with not found", bookId: "2", expectedStatus: http.StatusNotFound},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			req := httptest.NewRequest("GET", "/", nil)

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tc.bookId)

			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			w := httptest.NewRecorder()

			booksController.FindByID(w, req)

			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("Expected status %d, got %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func TestBooksControllerGetPage(t *testing.T) {
	booksController := controllers.NewBooksController(tests_mocks.BooksServiceMock{})
	testCases := []struct {
		name           string
		bookId         string
		page           string
		contentFormat  string
		expectedStatus int
	}{
		{name: "Test GetPage", bookId: "1", page: "1", contentFormat: "text", expectedStatus: http.StatusOK},
		{name: "Test GetPage with invalid id", bookId: "asdsa", page: "1", contentFormat: "text", expectedStatus: http.StatusBadRequest},
		{name: "Test GetPage with not found", bookId: "2", page: "1", contentFormat: "text", expectedStatus: http.StatusNotFound},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			req := httptest.NewRequest("GET", "/", nil)

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("bookid", tc.bookId)
			rctx.URLParams.Add("page", tc.page)
			rctx.URLParams.Add("contentFormat", tc.contentFormat)

			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			w := httptest.NewRecorder()

			booksController.GetPage(w, req)

			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("Expected status %d, got %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}
