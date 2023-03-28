package service

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"time"

	"http2/app/types"
	"http2/app/types/erors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (service *Service) GenToken(c *gin.Context, creds types.Credential) (string, error) {
	timeToDie, err := strconv.ParseInt(os.Getenv("TIME_TO_DIE"), 10, 64)
	if err != nil {
		return ":", err
	}
	claims := &types.Claims{
		Login: creds.Login,
		Password: creds.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(timeToDie) * time.Hour).Unix(),
		},
	}
	jwtKey, err := json.Marshal(os.Getenv("JWT_KEY"))
	if err != nil {
		return ":", err
	}

	encryption := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := encryption.SignedString(jwtKey)
	if err != nil {
		return ":", err
	}
	return jwtToken, nil
}

func (service *Service) ParseWithBearer(c *gin.Context, creds types.Credential) error{
	authorizationHeader := c.Request.Header.Get("authorization")
	if authorizationHeader == "" {
		return erors.NotFound
	}

	bearerToken := strings.Split(authorizationHeader, " ")

	if len(bearerToken) < 2 {
		return erors.NotFound
	}

	expiration, err := parseJWtToken(bearerToken[1])
	if err != nil {
		return erors.Invalid
		
	}

	for key, val := range *expiration {
		if key == "exp" {
			if time.Now().Unix() > int64(val.(float64)) {
				return erors.Invalid
			}
		}
	}
	for key, val := range *expiration {
		if key == "login" {
			if val != creds.Login {
				return erors.NotSame
			}
		}
	}
	return nil
}

func parseJWtToken(token string) (*jwt.MapClaims, error) {
	decod := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, decod, keyFunc)
	if err != nil {
		return nil, err
	}

	return &decod, nil
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, erors.Method
	}

	jwtKey, err := json.Marshal(os.Getenv("JWT_KEY"))
	if err != nil {
		return nil, err
	}

	return []byte(jwtKey), nil
}
