package types
import "github.com/golang-jwt/jwt"

type JWTUploadData struct {
	ID    uint64  `json:"id"`
	Login string `json:"login"`
	jwt.StandardClaims
}

// type Claims struct {
// 	Login    string `json:"login" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// 	jwt.StandardClaims
// }