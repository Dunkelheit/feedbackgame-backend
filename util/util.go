package util

import (
	"fmt"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
)

const secret = "626d1416bd941bc8a6eec67aed7cbd72a59dbdb6"

// FeedbackAppClaims are my custom claims
type FeedbackAppClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// StringToID converts a string to a ID
func StringToID(s string) (userID int64) {
	userID, _ = strconv.ParseInt(s, 10, 64)
	return
}

// EncodeToken encodes a token
func EncodeToken(id, email string) (string, error) {
	claims := FeedbackAppClaims{
		Username: id,
		Email:    email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// DecodeToken decodes a token
func DecodeToken(tokenString string) (string, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &FeedbackAppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return "", "", err
	}

	if claims, ok := token.Claims.(*FeedbackAppClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Username, claims.StandardClaims.ExpiresAt)
		return claims.Username, claims.Email, nil
	}

	return "", "", err
}
