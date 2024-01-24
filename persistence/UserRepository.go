package persistence

import "simple-go-web-boilerplate/models"

type UserRepository interface {
	CreateUser(string, string, int) error
	GetAllUsers() []models.User
	FindUserById(int) (models.User, error)
	DeleteUser(int) error
}
