package utils

import (
	"errors"
	"fmt"
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

func VerifyToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {

		// This checks if the signing method is a type of SigningMethodHMAC.
		// If its not ok, then execute the code in the if body.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println(ok)
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	//email := claims["email"].(string)
	userId := int64(claims["user_id"].(float64))

	return userId, nil
}
