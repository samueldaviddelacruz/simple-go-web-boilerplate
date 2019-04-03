package users

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/samueldaviddelacruz/simple-go-web-boilerplate/entities"
	"net/http"
)

type userRoute struct {
	userRepo entities.UserRepository
}

func (userRoute *userRoute) allUsers(w http.ResponseWriter, r *http.Request) {
	users := userRoute.userRepo.GetAllUsers()

	json.NewEncoder(w).Encode(users)
}

func New(userRepo entities.UserRepository) userRoute {
	return userRoute{userRepo}
}

func (userRoute *userRoute) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("", userRoute.allUsers)
	router.HandleFunc("/", userRoute.allUsers)
}
