package service

import jt "github.com/dgrijalva/jwt-go"

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
func ParseToken(k string) (uid int64, err error) {
	return 0, nil
}
