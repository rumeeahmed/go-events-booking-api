package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "loliamateapot"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{"email": email, "user_id": userId, "exp": time.Now().Add(time.Hour * 1).Unix()},
	)
	return token.SignedString([]byte(secretKey))
}
