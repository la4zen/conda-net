package model

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	jwt.StandardClaims
	ID int
}
