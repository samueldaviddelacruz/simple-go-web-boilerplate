package users

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	users := []User{
		{
			Firstname: "Samuel",
			Lastname:  "De la cruz",
			Age:       27,
		},
		{
			Firstname: "Oscar",
			Lastname:  "Martinez",
			Age:       26,
		},
	}

	json.NewEncoder(w).Encode(users)
}
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("", AllUsers)
	router.HandleFunc("/", AllUsers)
	//router.HandleFunc("/{title}", GetBook)
}
