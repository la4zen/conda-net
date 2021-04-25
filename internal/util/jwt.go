package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/la4zen/conda-net/internal/model"
)

var Key []byte = []byte("secretKeyy")

func CreateToken(user *model.User) (string, error) {
	claims := &model.Claims{
		ID: int(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(48 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(Key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
