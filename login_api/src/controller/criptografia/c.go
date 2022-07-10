package criptografia

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"login-api/src/model"
)

type Criptografia interface {
	Criptografa() string
	Compara(texto string) string
}

type CriptografiaMD5 struct {
	Texto string
}

func (c *CriptografiaMD5) Criptografa() string {
	hasher := md5.New()
	hasher.Write([]byte(c.Texto))
	texto := hex.EncodeToString(hasher.Sum(nil))
	return texto
}

func (c *CriptografiaMD5) Compara(senha string) bool {
	return senha == c.Criptografa()
}

func ValidateUser(u *model.User, r *model.Result) *model.Result {

	user := u
	result := r
	/*
		if user.isNewUser == true {
			result.isValid = false
			result.user = user

			return result
		}

		result.user = user
		result.isValid = true*/

	fmt.Println("{}", user)

	fmt.Println("{}", result)

	return result

}
