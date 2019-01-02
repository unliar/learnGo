package service

import (
	"errors"
	jt "github.com/dgrijalva/jwt-go"
)

type TokenPayload struct {
	UID         int64
	Status      int32
	PassUpdated int64
	jt.StandardClaims
}

// GeneratorToken 是用于生成token的方法
func GeneratorToken(t TokenPayload, k string) (s string, err error) {

	token := jt.NewWithClaims(jt.SigningMethodHS256, t)

	signedToken, err := token.SignedString(k)

	if err != nil {
		return "", err
	}
	return signedToken, nil

}

// ParseToken 解析token信息
func ParseToken(t string, k string) (uid int64, err error) {
	token, err := jt.ParseWithClaims(t, &TokenPayload{}, func(token *jt.Token) (i interface{}, e error) {
		return k, nil
	})

	if err != nil || !token.Valid {
		return 0, err
	}

	c, ok := token.Claims.(*TokenPayload)

	if !ok {
		return 0, errors.New("token.Claims not ok")
	}

	if c.Status != 1 {
		return 0, errors.New("user status unexpected")
	}

	return c.UID, nil
}
