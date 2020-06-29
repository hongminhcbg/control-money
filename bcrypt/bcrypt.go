package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptClient interface {
	GenerateFromPassword(password []byte, cost int) ([]byte, error)
	IsMatchingHashAndPassword(hashedPassword, password []byte) bool
}

type bcryptClientImpl struct {
}

func NewBcryptClient() BcryptClient {
	return &bcryptClientImpl{}
}

func (c *bcryptClientImpl) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}

func (c *bcryptClientImpl) IsMatchingHashAndPassword(hashedPassword, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	return err == nil
}
