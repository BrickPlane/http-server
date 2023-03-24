package service

import (
	"encoding/json"
	"os"
	"strings"
	"time"

	"http2/app/types/erors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (service *Service) ParseWithBearer(c *gin.Context) error{
	authorizationHeader := c.Request.Header.Get("authorization")
	if authorizationHeader == "" {
		return erors.NotFound
		// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Header is empty"})
	}
	
	bearerToken := strings.Split(authorizationHeader, " ")

	if len(bearerToken) < 2 {
		return erors.NotFound
		// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Header is empty"})
	}

	claims, err := parseJWtToken(bearerToken[1])
	if err != nil {
		return erors.Invalid
		// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Invalid token"})
	}

	for key, val := range *claims {
		if key == "exp" {
			if time.Now().Unix() > int64(val.(float64)) {
				return erors.Invalid
				// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Invalid token"})
			}
		}
	}
	// c.AbortWithStatusJSON(http.StatusOK, gin.H{"msg": "Token valid"})
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
		return nil, erors.Method
	}

	jwtKey, err := json.Marshal(os.Getenv("JWT_KEY"))
	if err != nil {
		return nil, err
	}

	return []byte(jwtKey), nil
}
