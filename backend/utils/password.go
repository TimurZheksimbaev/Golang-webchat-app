package utils
import "golang.org/x/crypto/bcrypt"

// Used for encrypting passwords and comparing them using bcrypt library

func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func ComparePasswords(first string, second string) error {
	err := bcrypt.CompareHashAndPassword([]byte(first), []byte(second))
	return err
}