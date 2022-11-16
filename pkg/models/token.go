package models

import jwt "github.com/golang-jwt/jwt"

type Token struct {
	UserID   uint
	UserName string
	Email    string
	*jwt.StandardClaims
}
