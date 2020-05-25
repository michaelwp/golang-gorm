package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPass(password []byte) ([]byte, error){
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {return nil, err}
	return hash, nil
}

func CompareHash(hash []byte, password []byte) error{
	err := bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {return err}
	return nil
}
