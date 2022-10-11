package Encrypt

import (
	"fmt"
	"testing"
)

func TestEncryptAES(t *testing.T) {
	// cipher key
	key := "thisis32bitlongpassphraseimusing" // 32位秘钥
	// plaintext
	pt := "This is a secret"
	c := EncryptAES([]byte(key), pt)
	// plaintext
	fmt.Println(pt)
	// ciphertext
	fmt.Println(c) // 145149d64a1a3c4025e67665001a3167
}

func TestDecryptAES(t *testing.T) {
	key := "thisis32bitlongpassphraseimusing"
	c := "145149d64a1a3c4025e67665001a3167"
	DecryptAES([]byte(key), c)
}
