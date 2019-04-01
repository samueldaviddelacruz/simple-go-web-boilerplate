package main

import (
	"net/http"
	"simple-web-boilerplate/routes/users"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(false)
	users.RegisterRoutes(router.PathPrefix("/users").Subrouter())

	http.ListenAndServe(":5000", router)
}
