package model

type User struct {
	Email     string
	Username  string
	Password  string
	IsAdmin   bool
	IsNewUser bool
}

func NewUser() *User {
	return &User{}
}
