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

func (service *Service) GenToken(c *gin.Context, creds types.User) (string, error) {
	timeToDie, err := strconv.ParseInt(os.Getenv("TIME_TO_DIE"), 10, 64)
	if err != nil {
		return ":", err
	}

	tokenData := &types.JWTUploadData{
		ID:    uint64(creds.ID),
		Login: creds.Login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(timeToDie) * time.Hour).Unix(),
		},
	}
	jwtKey, err := json.Marshal(os.Getenv("JWT_KEY"))
	if err != nil {
		return ":", err
	}

	encryption := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenData)
	jwtToken, err := encryption.SignedString(jwtKey)
	if err != nil {
		return ":", err
	}
	return jwtToken, nil
}

func (service *Service) ParseWithBearer(c *gin.Context) (*types.JWTUploadData, error) {
	authorizationHeader := c.Request.Header.Get("authorization")
	if authorizationHeader == "" {
		return nil, erors.NotFound
	}

	bearerToken := strings.Split(authorizationHeader, " ")

	if len(bearerToken) < 2 {
		return nil, erors.NotFound
	}

	tokenData, err := parseJWtToken(bearerToken[1])
	if err != nil {
		return nil, erors.Invalid

	}

	uploadData := &types.JWTUploadData{}
	for key, val := range *tokenData {
		switch key {
		case "exp":
			if time.Now().Unix() > int64(val.(float64)) {
				return nil, erors.Invalid
			}
		case "id":
			uploadData.ID = uint64(val.(float64))
		case "login":
			uploadData.Login = val.(string)
		default:
			return nil, err
		}
	}
	return uploadData, nil
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

func (service *Service) TokenVerification(tokenData *types.JWTUploadData) error {
	_, err := service.storage.GetUserByID(tokenData.ID)
	if err != nil {
		return err
	}
	return nil
}
