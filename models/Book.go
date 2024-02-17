package models

type Book struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Numpages int    `json:"numpages"`
}
