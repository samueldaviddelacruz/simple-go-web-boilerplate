package services_test

import (
	application_errors "simple-go-web-boilerplate/errors"
	"simple-go-web-boilerplate/services"
	tests_mocks "simple-go-web-boilerplate/tests"
	"testing"
)

func TestFindBook(t *testing.T) {
	// test the FindByID method
	booksService := services.NewBooksService(tests_mocks.BooksRepositoryMock{}, tests_mocks.PagesRepositoryMock{})
	testCases := []struct {
		name             string
		bookId           int
		expectedTitle    string
		expectedAuthor   string
		expectedNumpages int
		expErr           error
	}{
		{name: "Test FindByID", bookId: 1, expectedTitle: "test", expectedAuthor: "test", expectedNumpages: 1, expErr: nil},
		{name: "Test FindByID with invalid id", bookId: 2, expectedTitle: "", expectedAuthor: "", expectedNumpages: 0, expErr: application_errors.ErrBookNotFound},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			book, err := booksService.FindByID(tc.bookId)
			if tc.expErr != nil {
				if err != tc.expErr {
					t.Errorf("Expected error %v, got %v", tc.expErr, err)
				}
				return
			}
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if book.Title != tc.expectedTitle {
				t.Errorf("Expected title %s, got %s", tc.expectedTitle, book.Title)
			}
			if book.Author != tc.expectedAuthor {
				t.Errorf("Expected author %s, got %s", tc.expectedAuthor, book.Author)
			}
			if book.Numpages != tc.expectedNumpages {
				t.Errorf("Expected numpages %d, got %d", tc.expectedNumpages, book.Numpages)
			}
		})
	}
}

func TestGetPage(t *testing.T) {
	// test the GetPage method
	booksService := services.NewBooksService(tests_mocks.BooksRepositoryMock{}, tests_mocks.PagesRepositoryMock{})
	testCases := []struct {
		name            string
		bookId          int
		page            int
		contentFormat   string
		expectedContent string
		expErr          error
	}{
		{name: "Test GetPage", bookId: 1, page: 1, contentFormat: "text", expectedContent: "test", expErr: nil},
		{name: "Test GetPage with invalid book id", bookId: 2, page: 1, contentFormat: "text", expectedContent: "", expErr: application_errors.ErrPageNotFound},
		{name: "Test GetPage with invalid content format", bookId: 1, page: 1, contentFormat: "invalid", expectedContent: "", expErr: application_errors.ErrInvalidContentFormat},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			content, err := booksService.GetPage(tc.bookId, tc.page, tc.contentFormat)
			if tc.expErr != nil {
				if err != tc.expErr {
					t.Errorf("Expected error %v, got %v", tc.expErr, err)
				}
				return
			}
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if content != tc.expectedContent {
				t.Errorf("Expected content %s, got %s", tc.expectedContent, content)
			}
		})
	}
}
