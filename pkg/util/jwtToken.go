package util

import (
	"errors"
	"fmt"

	"os"
	"time"

	"github.com/aparnasukesh/shoezone/pkg/domain"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(userData domain.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["role"] = userData.Isadmin
	claims["email"] = userData.Email
	claims["userid"] = userData.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()

	secretKey := []byte(os.Getenv("JWT_secret_key"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_secret_key")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func GetRole(token *jwt.Token) (interface{}, error) {
	claims, _ := token.Claims.(jwt.MapClaims)
	role := claims["role"]
	return role, nil
}

func GetUserID(token *jwt.Token) (int, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims type")
	}

	userID, ok := claims["userid"].(float64)
	if !ok {
		return 0, errors.New("user ID not found or not a number in claims")
	}

	return int(userID), nil
}
