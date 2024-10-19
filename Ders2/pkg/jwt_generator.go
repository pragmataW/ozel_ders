package pkg

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtGenerator struct {
	Key string
}

func (j *JwtGenerator)Generate(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp": time.Now().Add(time.Hour * 24),
	}

	signer := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := signer.SignedString([]byte(j.Key))
	if err != nil{
		return "", err
	}
	return token, nil
}

func NewJwtGenerator(key string) JwtGenerator {
	return JwtGenerator{
		Key: key,
	}
}