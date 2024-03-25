package common

import (
	"ginEssential/model"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("a_secret_key")

type Claims struct {
	UserID uint
	jwt.RegisteredClaims
}

func ReleaseClaims(user model.User) (string, error) {
	claims := &Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			Issuer:    "oceanlearn.tech",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "user_token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
