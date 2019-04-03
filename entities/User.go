package entities

type User struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

type UserRepository interface {
	CreateUser(string, string, int)
	GetAllUsers() []User
	FindUserById(int) User
	DeleteUser(int)
}
