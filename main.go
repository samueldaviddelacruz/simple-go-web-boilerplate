package main

import (
	"context"
	"encoding/json"
	"net/http"
	"simple-go-web-boilerplate/models"
	"simple-go-web-boilerplate/persistence"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type contextKey string

const userKey contextKey = "user"

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		hello([]string{
			"Samuel",
		}).Render(r.Context(), w)
	})
	r.Route("/users", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			userRepo := persistence.NewInMemoryUserRepository()
			users := userRepo.GetAllUsers()
			json.NewEncoder(w).Encode(users)
		})

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			userRepo := persistence.NewInMemoryUserRepository()
			var user models.User
			json.NewDecoder(r.Body).Decode(&user)
			userRepo.CreateUser(user.Firstname, user.Lastname, user.Age)
			w.Write([]byte("create user"))
		})

		r.Route("/{userId}", func(r chi.Router) {
			r.Use(UserContext)
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				user, ok := r.Context().Value(userKey).(models.User)
				if !ok {
					http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
					return
				}
				json.NewEncoder(w).Encode(user)
			})
			r.Put("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("update user"))
			})
			r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("delete user"))
			})
		})
	})
	http.ListenAndServe(":5000", r)
}
func UserContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		userRepo := persistence.NewInMemoryUserRepository()
		userId := chi.URLParam(r, "userId")
		userIdInt, err := strconv.Atoi(userId)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		user, err := userRepo.FindUserById(userIdInt)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), userKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
