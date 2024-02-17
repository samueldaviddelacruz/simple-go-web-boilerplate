# Golang REST API

This is a simple Golang REST API boilerplate. The API provides endpoints to retrieve information about books stored in an SQLite3 database, it can serve as base project and extended further.

## Requirements
- Go 1.13 or later (for module support)

## Installation and running
1. Install packages using `go get -u ./...`

2. Run the application:
 ```go run main.go```

The application will be running on port 3000.

## Available Routes

### 1. Get List of Books
- **Endpoint:** `/books`
- **Method:** `GET`
- **Example:**

```curl http://localhost:3000/books```

### 2. Get a Book by ID
- **Endpoint:** `/books/{id}`
- **Method:** `GET`
- **Example:**

```curl http://localhost:3000/books/1```
### 3. Get a Page of a Book
- **Endpoint:** `/books/{bookid}/pages/{page}/{contentFormat}`
- **Method:** `GET`
- **Parameters:**
- `{bookid}`: ID of the book
- `{page}`: Page number
- `{contentFormat}`: Content format (text or html)
- **Example:**

```curl http://localhost:3000/books/1/pages/3/text```


## Data
An SQLite database named `books.db` is provided inside the `./data` folder. It contains 5 books, each with multiple pages.

## Running Tests
To run tests, use the following command:
```go test -v ./...```