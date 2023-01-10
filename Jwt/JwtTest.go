package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyCustomClaims struct {
	ID       int64
	Username string
	jwt.StandardClaims
}
type StandardClaims struct {
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`
}

func GenerateToken() (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(300 * time.Second)
	issuer := "frank"
	claims := MyCustomClaims{
		ID:       10001,
		Username: "frank",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("golang"))
	return token, err
}
func main() {

	fmt.Printf(GenerateToken())
}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MTAwMDEsIlVzZXJuYW1lIjoiZnJhbmsiLCJleHAiOjE2NjkwMzUyNDgsImlzcyI6ImZyYW5rIn0.KlrQo0nxTOSd48y169W0F57ij6ypadG5vRInLv5X8zg%!(EXTRA <nil>)
