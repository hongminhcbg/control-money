package middlewares

import (
	"time"

	jwtLib "github.com/dgrijalva/jwt-go"
)

type JWT interface {
	CreateToken(id int64) (string, error)
}

type jwtImpl struct {
	secretKey string
}

func NewJWT(secretKey string) JWT {
	return &jwtImpl{secretKey: secretKey}
}

func (c *jwtImpl) CreateToken(id int64) (string, error) {
	token := jwtLib.New(jwtLib.GetSigningMethod("HS256"))

	token.Claims = jwtLib.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	if tokenString, err := token.SignedString([]byte(c.secretKey)); err != nil {
		return "", nil
	} else {
		return tokenString, nil
	}
}
