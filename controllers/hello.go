package controllers

import (
	"net/http"
	"simple-go-web-boilerplate/views"
)

type helloController struct {
}

func (h helloController) Index(w http.ResponseWriter, r *http.Request) {
	views.Render(views.Hello("Hi"))(w, r)
}

func NewHelloController() helloController {
	return helloController{}
}
