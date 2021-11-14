package Modules

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"hash"
	"io"
	"os"
)

//calculates the signature of hashed using RSASSA-PKCS1-V1_5-SIGN from RSA PKCS#1 v1.5.
func SignPKCS1v15(plaintext string, privKey rsa.PrivateKey) string {
	// crypto/rand.Reader is a good source of entropy for blinding the RSA
	// operation.
	rng := rand.Reader
	hashed := sha256.Sum256([]byte(plaintext))
	signature, err := rsa.SignPKCS1v15(rng, &privKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
		return "Error from signing"
	}
	return base64.StdEncoding.EncodeToString(signature)
}

// ECDSA

func SignECDSA(plaintext string, privateKey *ecdsa.PrivateKey) string {

	// Sign ecdsa style
	var h hash.Hash = md5.New()

	io.WriteString(h, plaintext)
	signhash := h.Sum(nil)

	r, s, serr := ecdsa.Sign(rand.Reader, privateKey, signhash)
	Check_erorr(serr)

	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)
	return string(signature)
}
