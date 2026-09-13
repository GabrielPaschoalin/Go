package utils

import "golang.org/x/crypto/bcrypt"

// FIle to hash password

// HashPassword gera o hash bcrypt de uma senha em texto plano.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

// CheckPasswordHash confere se a senha em texto plano corresponde ao hash salvo.
func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
