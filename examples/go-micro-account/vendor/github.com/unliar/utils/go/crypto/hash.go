package crypto

import "golang.org/x/crypto/bcrypt"

func HashString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), 10)
	return string(b), err
}

func MatchHashString(s, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(s))
	return err == nil
}
