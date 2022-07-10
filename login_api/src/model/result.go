package model

type Result struct {
	user    User //`json:`
	isValid bool //`json:`
}

func NewResult() *Result {
	return &Result{}
}
