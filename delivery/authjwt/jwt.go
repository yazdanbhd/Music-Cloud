package authjwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type Token struct {
	secretKey  []byte
	signMethod jwt.SigningMethod
}

func New(secretKey []byte, signMethod jwt.SigningMethod) Token {
	return Token{secretKey: secretKey, signMethod: signMethod}
}

func (t Token) CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(t.signMethod,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(t.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (t Token) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return t.secretKey, nil
	})

	if err != nil || !token.Valid {
		return jwt.MapClaims{}, fmt.Errorf("invalid token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("could not parse claims")
}
