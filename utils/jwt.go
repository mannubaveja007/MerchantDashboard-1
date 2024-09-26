package utils

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

var secretKey = []byte("your_secret_key")

func GenerateJWT(username string) (string, error) {
    claims := &jwt.StandardClaims{
        Subject: username,
        ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // Token expiration
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secretKey)
}
