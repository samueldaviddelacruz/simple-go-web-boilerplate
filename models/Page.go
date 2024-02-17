package models

type Page struct {
	Id          int    `json:"id"`
	Bookid      int    `json:"bookid"`
	ContentTxt  string `json:"content"`
	ContentHTML string `json:"contenthtml"`
	Number      int    `json:"number"`
}
