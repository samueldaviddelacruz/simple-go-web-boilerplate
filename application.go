package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/samueldaviddelacruz/simple-go-web-boilerplate/Persistence"
	"github.com/samueldaviddelacruz/simple-go-web-boilerplate/routes/users"
	"log"
	"net/http"
)

func main() {
	portFlag := flag.Int("port", 5000, "the port to which the server will listen to")
	router := mux.NewRouter().StrictSlash(false)
	usersDB := &Persistence.InMemoryUsersDB{}

	usersRoute := users.New(usersDB)
	usersSubRouter := router.PathPrefix("/users").Subrouter()
	usersRoute.RegisterRoutes(usersSubRouter)

	address := fmt.Sprintf(":%d", *portFlag) //":5000
	fmt.Printf("Listening on port %d\n", *portFlag)
	log.Fatal(http.ListenAndServe(address, router))
}
