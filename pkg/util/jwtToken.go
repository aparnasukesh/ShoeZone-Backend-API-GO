package util

import (
	"fmt"
	"os"
	"time"

	"github.com/aparnasukesh/shoezone/pkg/domain"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(userData domain.User) (string, error) {
	// Create a new token and specify the signing method
	token := jwt.New(jwt.SigningMethodHS256)

	// Create a map to store the claims
	claims := token.Claims.(jwt.MapClaims)

	// Set the claims (e.g., username, expiration, issued at, etc.)

	claims["role"] = userData.Isadmin
	claims["email"] = userData.Email
	claims["userid"] = userData.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time (e.g., 24 hours from now)
	claims["iat"] = time.Now().Unix()                     // Token issued at

	// Sign the token with a secret key
	secretKey := []byte(os.Getenv("JWT_secret_key"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	// Parse the token with the provided secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key used for verification
		return []byte(os.Getenv("JWT_secret_key")), nil
	})

	// Check for parsing errors
	if err != nil {
		return nil, err
	}

	return token, nil
}

func Getrole(token *jwt.Token) (interface{}, error) {
	claims, _ := token.Claims.(jwt.MapClaims)
	// if !ok {
	// 	return false, errors.New("Error extracting claims")
	// }
	role := claims["role"]
	return role, nil

	// if role := claims["role"]; role == true {
	// 	return true, nil
	// }
	// return false, nil
}
