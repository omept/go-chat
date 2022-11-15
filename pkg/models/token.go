package models

import jwt "github.com/golang-jwt/jwt"

type Token struct {
	UserID uint
	Uname  string
	Email  string
	*jwt.StandardClaims
}
