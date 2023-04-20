package types
import "github.com/golang-jwt/jwt"

type JWTUploadData struct {
	ID    uint64  `json:"id"`
	Login string `json:"login"`
	jwt.StandardClaims
}