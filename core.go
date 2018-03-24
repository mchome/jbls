package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"io/ioutil"
	"log"
	"os"
)

func sign(message string, key []byte) string {
	block, _ := pem.Decode([]byte(key))
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	hash := md5.New()
	hash.Write([]byte(message))
	digest := hash.Sum(nil)
	sign, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.MD5, digest)
	if err != nil {
		log.Print(err)
		return ""
	}
	signature := hex.EncodeToString(sign)
	return signature
}

func validateKey(filepath string) (bool, []byte) {
	file, _ := os.Open(filepath)
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	block, _ := pem.Decode([]byte(data))
	if block == nil {
		log.Print("Failed to parse pem block.")
		return false, nil
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if priv == nil {
		log.Print("Failed to parse PKCS1 private key.")
		return false, nil
	}

	return true, data
}
