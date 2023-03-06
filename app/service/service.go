package service

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"http2/app/storage"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)


type Service struct{
	// structure storage.SignKey
}

func NewService() *Service {
	return new(Service)
}

func (s *Service) SignToken(c *gin.Context, creds storage.Credential) (string, error) {
	timeToDie, err := strconv.ParseInt(os.Getenv("TIME_TO_DIE"), 10, 64)
	if err != nil {
		return "", err
	}
	claims := &storage.Claims{
		Login:    creds.Login,
		Password: creds.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(timeToDie) * time.Hour).Unix(),
		},
	}

	jwtKey, err := json.Marshal(os.Getenv("JWT_KEY"))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSignStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenSignStr, nil
}
