package model

type User struct {
	email    string //`json:""`
	username string //`json:""`
	password string //`json:""`
	isAdmin  bool   //`json:""`
}