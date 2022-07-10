package model

type Result struct {
	User    *User
	IsValid bool
}

func NewResult() *Result {
	return &Result{}
}
