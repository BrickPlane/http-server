package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"http2/app/storage"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type IStorage interface {
	StorageIn(c *gin.Context, data string) error
	// ParseWithBearer(c *gin.Context)
}
type Service struct{
	storage IStorage
}

func NewService(storage IStorage) *Service {
	return &Service{
		storage: storage,
	}
}

func (service *Service) SignToken(c *gin.Context, creds storage.Credential) (string, error) {
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
	errr := service.storage.StorageIn(c , tokenSignStr) 
	if errr != nil {
		return "", err
	}
	return tokenSignStr, nil
}

func (service *Service)ParseWithBearer(c *gin.Context) {
	authorizationHeader := c.Request.Header.Get("authorization")
	if authorizationHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg" : "Header is empty"})
	}

	bearerToken := strings.Split(authorizationHeader, " ")

	if len(bearerToken) < 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg" : "Header is empty"})
	}

	claims, err := parseJWtToken(bearerToken[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg" : "Invalid token"})
	}

	for key, val := range *claims {
		if key == "exp" {
			if time.Now().Unix() > int64(val.(float64)) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg" : "Invalid token"}) 
			}
		}
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"msg" : "Token valid"}) 
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