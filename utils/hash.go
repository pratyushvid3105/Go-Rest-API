package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error){
	// []byte(password) converts the password string into a byte Slice. In addition, GenerateFromPassword needs a cost parameter, which simply controls how complex the hashing will be. And here at the moment, a value of 14 should give us a secure hashed password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}