package handlers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
)

type Session struct {
	PrivateKey string
	PublicKey  string `json:"publicKey"`
	UserId     string `json:"userId"`
}

func ExchangeKeys(w http.ResponseWriter, r *http.Request) {
	privateKey, publicKey := generateKeys()
	w.Write([]byte(privateKey))
	w.Write([]byte(publicKey))
}

func generateKeys() (string, string) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		fmt.Printf("Cannot generate RSA key\n")
		os.Exit(1)
	}
	publicKey := &privateKey.PublicKey

	// convert keys to byte array
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Printf("Cannot convert public key to byte array\n")
		os.Exit(1)
	}
	return string(privateKeyBytes), string(publicKeyBytes)
}
