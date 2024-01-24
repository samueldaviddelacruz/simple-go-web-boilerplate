package persistence

import (
	"errors"
	"simple-go-web-boilerplate/models"
)

var users = []models.User{
	{
		Id:        1,
		Firstname: "Samuel",
		Lastname:  "De la cruz",
		Age:       27,
	},
	{
		Id:        2,
		Firstname: "Oscar",
		Lastname:  "Martinez",
		Age:       26,
	},
}

type inMemoryUserRepository struct{}

func NewInMemoryUserRepository() UserRepository {
	return &inMemoryUserRepository{}
}
func (inMemoryUserDB *inMemoryUserRepository) CreateUser(firstName string, lastName string, age int) error {
	users = append(users, models.User{Firstname: firstName, Lastname: lastName, Age: age})
	return nil
}
func (inMemoryUserDB *inMemoryUserRepository) GetAllUsers() []models.User {
	return users
}

func (inMemoryUserDB *inMemoryUserRepository) FindUserById(userId int) (models.User, error) {
	for index, user := range users {
		if user.Id == userId {
			return users[index], nil
		}
	}
	return models.User{}, errors.New("user not found")
}
func (inMemoryUserDB *inMemoryUserRepository) DeleteUser(userId int) error {
	return nil
}
