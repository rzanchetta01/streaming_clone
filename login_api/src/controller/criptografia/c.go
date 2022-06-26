package criptografia

import (
	"crypto/md5"
	"encoding/hex"
)

type Criptografia interface {
	Criptografa() string
	Compara(texto string) string
}

type CriptografiaMD5 struct {
	Texto string
}

func (c CriptografiaMD5) Criptografa() string {
	hasher := md5.New()
	hasher.Write([]byte(c.Texto))
	texto := hex.EncodeToString(hasher.Sum(nil))
	return texto
}

func (c CriptografiaMD5) Compara(senha string) bool {
	return senha == c.Criptografa()
}
