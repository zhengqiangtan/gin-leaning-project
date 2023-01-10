package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/cidertool/asc-go/asc"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const issuerID = "10696be8-e055-47c7-af84-d934999b7abf"
const keyID = "T22Y827GNG"

// const keyFile = "C:\\Users\\HNK7WC3\\Desktop\\appstore\\SubscriptionKey_Z364X8S2L2.p8"
const keyFile = "C:\\Users\\HNK7WC3\\Downloads\\AuthKey_T22Y827GNG.p8"

/*
*
GO111MODULE=on
go get github.com/gbrlsnchs/jwt/v3
*/
func main() {

	//privateKey, err := privateKeyFromFile()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//authToken, err := generateAuthToken(privateKey)
	//if err != nil {
	//	log.Fatal(err)
	//}

	authToken, _ := generateAuthToken()

	fmt.Println(authToken)

	client := &http.Client{}

	qs := url.Values{}
	qs.Set("fields[customerReviews]", "body")
	qs.Set("limit", "10")

	jsonBody := []byte(`{"fields[customerReviews]": "body","limit":2}`)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(
		http.MethodGet,
		//"https://api.appstoreconnect.apple.com/v1/apps/1468073139/customerReviews?"+qs.Encode(),
		"https://api.appstoreconnect.apple.com/v1/apps/1468073139/customerReviews",
		bodyReader,
	)
	if err != nil {
		log.Fatal(err)
	}

	//authToken := "eyJhbGciOiJFUzI1NiIsImtpZCI6IlQyMlk4MjdHTkciLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiIxMDY5NmJlOC1lMDU1LTQ3YzctYWY4NC1kOTM0OTk5YjdhYmYiLCJleHAiOjE2NjkxMDE0NTksImF1ZCI6ImFwcHN0b3JlY29ubmVjdC12MSJ9.u0znDNKsmTuesNSUg7VflwdDWH1t_Jd-F6JsIs1r17dAqCylwzTaCcEjhqFInBs2czbw92x8QLiZa2sX4NT98w"
	//authToken = "eyJhbGciOiJFUzI1NiIsImtpZCI6IlQyMlk4MjdHTkciLCJ0eXAiOiJKV1QifQ.eyJhdWQiOlsiYXBwc3RvcmVjb25uZWN0LXYxIl0sImV4cCI6MTY2OTI1NzQ2OS43MzI3ODEyLCJpc3MiOiIxMDY5NmJlOC1lMDU1LTQ3YzctYWY4NC1kOTM0OTk5YjdhYmYifQ.YixyjWCAJWGCppOexWOXRDOqjLd76MBL6DlJthoST__UnRMerqFq6_quBUjpYkIyabfo4p-qVjoUlmOHsRdqjw"
	//authToken= "eyJhbGciOiJFUzI1NiIsImtpZCI6IlQyMlk4MjdHTkciLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJhcHBzdG9yZWNvbm5lY3QtdjEiLCJleHAiOjE2NjkyNjE0NTIsImlzcyI6IjEwNjk2YmU4LWUwNTUtNDdjNy1hZjg0LWQ5MzQ5OTliN2FiZiJ9.JW0NSLu-tTHtE69YGxknJTNgPvpcKVVbTNeLVhH8KoruI9-Ss29ZXp1uxBkLU7BH5a_zo4z-QR607jm71zf0wA"
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("User-Agent", "App Store Connect Client")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bytes))

}

func privateKeyFromFile() (*ecdsa.PrivateKey, error) {

	bytes, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, errors.New("AuthKey must be a valid .p8 PEM file")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pk := key.(type) {
	case *ecdsa.PrivateKey:
		return pk, nil
	default:
		return nil, errors.New("AuthKey must be of type ecdsa.PrivateKey")
	}

}

type standardJWTGenerator struct {
	keyID          string
	issuerID       string
	expireDuration time.Duration
	privateKey     *ecdsa.PrivateKey

	token string
}

func parsePrivateKey(blob []byte) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode(blob)
	if block == nil {
		return nil, nil
	}

	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	if key, ok := parsedKey.(*ecdsa.PrivateKey); ok {
		return key, nil
	}

	return nil, nil
}

func generateAuthToken() (string, error) {

	//expirationTimestamp := time.Now().Add(30 * time.Minute)

	expiryDuration := 20 * time.Minute

	privateKey1, _ := os.ReadFile("C:\\Users\\HNK7WC3\\Downloads\\AuthKey_T22Y827GNG.p8")

	privateKey, err := parsePrivateKey(privateKey1)
	if err != nil {
		log.Fatal(err)
	}

	gen := &standardJWTGenerator{
		keyID:          keyID,
		issuerID:       issuerID,
		privateKey:     privateKey,
		expireDuration: expiryDuration,
	}
	token, err := gen.Token()

	//token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
	//	"iss": issuerID,
	//	"exp": expirationTimestamp.Unix(),
	//	"aud": "appstoreconnect-v1",
	//})

	//token := jwt.NewWithClaims(jwt.SigningMethodES256,jwt.StandardClaims{
	//	Audience:  "appstoreconnect-v1",
	//	Issuer:    issuerID,
	//	ExpiresAt: expirationTimestamp.Unix(),
	//})
	//
	//token.Header["kid"] = keyID
	//
	//tokenString, err := token.SignedString(privateKey)
	//if err != nil {
	//	return "", err
	//}

	return token, nil

}

func (g *standardJWTGenerator) Token() (string, error) {
	//if g.IsValid() {
	//	return g.token, nil
	//}

	t := jwt.NewWithClaims(jwt.SigningMethodES256, g.claims())
	t.Header["kid"] = g.keyID

	token, err := t.SignedString(g.privateKey)
	if err != nil {
		return "", err
	}

	g.token = token

	return token, nil
}

func GenToken() (interface{}, error) {
	secret, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}
	auth, err := asc.NewTokenConfig(keyID, issuerID, 20*time.Minute, secret)
	fmt.Println(auth.Client())

	return nil, nil
}

func (g *standardJWTGenerator) claims() jwt.Claims {
	//expiry := time.Now().Add(g.expireDuration)
	expirationTimestamp := time.Now().Add(30 * time.Minute)

	return jwt.StandardClaims{
		//Audience:  jwt.ClaimStrings{"appstoreconnect-v1"},
		Audience:  "appstoreconnect-v1",
		Issuer:    g.issuerID,
		ExpiresAt: expirationTimestamp.Unix(),
	}
}

//func genToken2() (interface{},interface{}) {
//	expirationTimestamp := time.Now().Add(30 * time.Minute)
//
//
//
//
//	t := jwt.NewWithClaims(jwt.SigningMethodES256, 	 jwt.StandardClaims{
//		Audience:  jwt.ClaimStrings{"appstoreconnect-v1"},
//		Issuer:   issuerID,
//		ExpiresAt: expirationTimestamp
//	})
//	t.Header["kid"] = g.keyID
//
//	token, err := t.SignedString(g.privateKey)
//	if err != nil {
//		return "", err
//	}
//}

//func main() {
//	GenToken()
//}
