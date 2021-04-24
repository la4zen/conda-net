package util

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/la4zen/conda-net/internal/model"
)

var key []byte = []byte("secretKeyy")

func CreateToken(user *model.User) (string, error) {
	claims := &model.Claims{
		ID: int(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(48 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(user *model.User, tokenHeader string) error {
	tokenString := strings.Split(tokenHeader, " ")[1]
	claims := &model.Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return err
	}
	user.ID = uint(claims.ID)
	return nil
}
