package user_service

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"http2/app/storage"
	"http2/app/types"
	"http2/app/types/userDB"

	"http2/app/types/errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
)

func (service *Service) GenToken(c *gin.Context, creds user_types.User) (*string, error) {
	timeToDie, err := strconv.ParseInt(os.Getenv("TIME_TO_DIE"), 10, 64)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	encryption := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenData)
	jwtToken, err := encryption.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}
	err = SaveToken(jwtToken, creds.Login, 10)
	if err != nil {
		return nil, err
	}

	return &jwtToken, nil
}

func SaveToken(token, login string, expires int64) error {
	rdb, ctx := storage.RedisDB()
	_, err := rdb.SetNX(ctx, login, token, time.Duration(expires)*time.Hour).Result()
	if err != nil {
		return err
	}
	return nil
}

func (service *Service) ParseWithBearer(c *gin.Context) (*types.JWTUploadData, error) {
	authorizationHeader := c.Request.Header.Get("authorization")
	if authorizationHeader == "" {
		return nil, errors.NotFound
	}

	bearerToken := strings.Split(authorizationHeader, " ")

	if len(bearerToken) < 2 {
		return nil, errors.NotFound
	}

	tokenData, err := parseJWtToken(bearerToken[1])
	if err != nil {
		return nil, errors.Invalid

	}

	uploadData := &types.JWTUploadData{}
	for key, val := range *tokenData {
		switch key {
		case "exp":
			if time.Now().Unix() > int64(val.(float64)) {
				return nil, errors.Invalid
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
		return nil, errors.Method
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

func (service *Service) CheckToken(login string, c *gin.Context) error {
	authorizationHeader := c.Request.Header.Get("authorization")
	if authorizationHeader == "" {
		return errors.NotFound
	}

	bearerToken := strings.Split(authorizationHeader, " ")

	if len(bearerToken) < 2 {
		return errors.NotFound
	}

	rdb, ctx := storage.RedisDB()
	val, err := rdb.Get(ctx, login).Result()
	if err == redis.Nil {
		return err
	} else if err != nil {
		return err
	}
	
	if bearerToken[1] != val {
		fmt.Println("error", err)
		return err
	}
	fmt.Println("ok")
	return nil
}
