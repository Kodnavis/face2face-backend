package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ErrInvalidToken = errors.New("invalid token")

func GenerateToken(login string, token_lifespan int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": login,
		"exp": time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix(),
	})

	token_string, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return token_string, nil
}

func ExtractToken(token_string string, jwt_secret string) (string, error) {
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwt_secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return "", ErrInvalidToken
		}

		user_login := claims["sub"].(string)
		return user_login, nil
	}

	return "", ErrInvalidToken
}
