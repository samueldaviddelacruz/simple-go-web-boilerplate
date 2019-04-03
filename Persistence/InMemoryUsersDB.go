package Persistence

import "github.com/samueldaviddelacruz/simple-go-web-boilerplate/entities"

type InMemoryUsersDB struct {
}

var users = []entities.User{
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

func (inMemoryUserDB *InMemoryUsersDB) CreateUser(firstName string, lastName string, age int) {
	users = append(users, entities.User{Firstname: firstName, Lastname: lastName, Age: age})
}
func (inMemoryUserDB *InMemoryUsersDB) GetAllUsers() []entities.User {
	return users
}

func (inMemoryUserDB *InMemoryUsersDB) FindUserById(userId int) entities.User {
	for index, user := range users {
		if user.Id == userId {
			return users[index]
		}
	}
	return entities.User{}
}
func (inMemoryUserDB *InMemoryUsersDB) DeleteUser(userId int) {

}
