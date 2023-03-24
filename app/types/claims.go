package types

import "github.com/golang-jwt/jwt"


type Claims struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	jwt.StandardClaims
}
