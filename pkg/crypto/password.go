package crypto

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Password interface {
	Salt(password string) string
	Compare(hashedPwd string, plainPwd []byte) bool
}

type CryptoPassword struct{}

func (p *CryptoPassword) Salt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Fatal("Error to salt password")
	}

	return string(hash)
}

func (p *CryptoPassword) Compare(hashedPwd string, plainPwd []byte) bool {
	return false
}
