package service

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"http2/app/storage"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func SignToken(c *gin.Context, creds storage.Credential) (string, error) {
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

func Parse(c *gin.Context, tokenSignStr string) (*jwt.Token, error) {
	clam := jwt.MapClaims{}

	jwtKey, err := json.Marshal(os.Getenv("JWT_KEY"))
	if err != nil {
		return nil, err
	}
	tok, err := jwt.ParseWithClaims(
		tokenSignStr,
		clam,
		func(tok *jwt.Token) (interface{}, error) { return []byte(jwtKey), nil },
	)
	if err != nil {
		return nil, err
	}

	for key, val := range clam {
		if key == "exp" {
			expiredDate := val
			if time.Now().Unix() > int64(expiredDate.(float64)) {
				return nil, errors.New("token expired")
			}
		}
	}

	timeToDie, err := strconv.ParseInt(os.Getenv("TIME_TO_DIE"), 10, 64)
	if err != nil {
		return nil, err
	}

	clam["exp"] = timeToDie
	return tok, nil
}

func WithBearer(c *gin.Context) error {

	authorizationHeader := c.Request.Header.Get("authorization")
	if authorizationHeader == "" {
		return errors.New("Header is empty")
	}

	bearerToken := strings.Split(authorizationHeader, " ")

	if len(bearerToken) < 2 {
		return errors.New("Header is empty")
	}

	claims, err := parseJWtToken(bearerToken[1])
	if err != nil {
		return err
	}

	for key, val := range *claims {
		if key == "exp" {
			if time.Now().Unix() > int64(val.(float64)) {
				return errors.New("Invalid token")
			}
		}
	}

	return nil
}

func parseJWtToken(token string) (*jwt.MapClaims, error) {
	clam := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, clam, keyFunc)
	if err != nil {
		return nil, err
	}

	return &clam, nil
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("method error")
	}

	jwtKey, err := json.Marshal(os.Getenv("JWT_KEY"))
	if err != nil {
		return nil, err
	}

	return []byte(jwtKey), nil
}