package main

import (
	"crypto/x509"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"time"
)

const issuerID1 = "10696be8-e055-47c7-af84-d934999b7abf"
const keyID1 = "T22Y827GNG"

// const keyFile = "C:\\Users\\HNK7WC3\\Desktop\\appstore\\SubscriptionKey_Z364X8S2L2.p8"
const keyFile1 = "C:\\Users\\HNK7WC3\\Downloads\\AuthKey_T22Y827GNG.p8"

func main() {
	token := &jwt.Token{
		Header: map[string]interface{}{
			"typ": "JWT",
			"kid": keyID1,
			"alg": jwt.SigningMethodES256.Alg(),
		},
		Claims: jwt.MapClaims{
			"iss": issuerID1,
			"iat": time.Now().Unix(),
			"exp": time.Now().Add(3600 * time.Second).Unix(),
			"aud": "appstoreconnect-v1",
		},
		Method: jwt.SigningMethodES256,
	}

	privatePem, err := ioutil.ReadFile(keyFile1)
	ecdsaKey, err := x509.ParsePKCS8PrivateKey(privatePem)
	if err != nil {
		//t.Log("ecdsaKey Error...", err)
		return
	}
	tk, err := token.SignedString(ecdsaKey)
	fmt.Println(tk, err)
}
