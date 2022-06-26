package main

import (
	"fmt"
	"login-api/src/controller/criptografia"
)

func init() {}

func main() {
	var c = criptografia.CriptografiaMD5{
		Texto: "Teste",
	}

	fmt.Println(c.Criptografa())
}
