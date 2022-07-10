package model

type User struct {
	email     string
	username  string
	password  string
	isAdmin   bool
	isNewUser bool
}

func NewUser() *User {
	return &User{}
}
