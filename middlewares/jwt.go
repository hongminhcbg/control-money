package middlewares

import (
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
)

type JWT interface {
	CreateTocken(id int64) (string, error)
}

type jwtImpl struct {
	secretKey string
}

func NewJWT(secretKey string) JWT {
	return &jwtImpl{secretKey: secretKey}
}

func (c *jwtImpl) CreateTocken(id int64) (string, error) {
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))

	token.Claims = jwt_lib.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	if tokenString, err := token.SignedString([]byte(c.secretKey)); err != nil {
		return "", nil
	} else {
		return tokenString, nil
	}
}
