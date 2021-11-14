package Modules

import (
	"bufio"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

//Takes the HOME path from the environment varibles
var PrivateKeyPath = os.Getenv("HOME") + "/PrivateKey.pem"

//Takes the HOME path from the environment varibles
var PublicKeyPath = os.Getenv("HOME") + "/PublicKey.pem"

//Check if file exists and if it is a directory
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//Check if erorr exists
func Check_erorr(e error) {
	if e != nil {
		panic(e)
	}
}

// Method to store the RSA keys in pkcs1 Format
func SavePKCS1RSAPublicPEMKey(fName string, pubkey *rsa.PublicKey) {
	//converts an RSA public key to PKCS#1, ASN.1 DER form.
	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(pubkey),
	}
	pemfile, err := os.Create(fName)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
	defer pemfile.Close()
	err = pem.Encode(pemfile, pemkey)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

// Method to store the RSA keys in pkcs8 Format
func SavePKCS8RSAPEMKey(fName string, key *rsa.PrivateKey) {

	fmt.Println(fName)
	outFile, err := os.Create(fName)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
	defer outFile.Close()
	//converts a private key to ASN.1 DER encoded form.
	var privateKey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	err = pem.Encode(outFile, privateKey)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

//Load RSA private key which is stored in PEM format and create RSA private Key using go x509.ParsePKCS1PrivateKey
func LoadRSAPrivatePemKey(fileName string) *rsa.PrivateKey {
	privateKeyFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
	data, _ := pem.Decode([]byte(pembytes))
	privateKeyFile.Close()
	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
	return privateKeyImported
}

//generate and save keypair
func Save_keypair() {

	if !FileExists(PublicKeyPath) {

		// Generate Alice RSA keys Of 2048 Bits
		PrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// Extract Public Key from RSA Private Key
		PublicKey := PrivateKey.PublicKey

		SavePKCS8RSAPEMKey(PrivateKeyPath, PrivateKey)
		SavePKCS1RSAPublicPEMKey(PublicKeyPath, &PublicKey)
	}
}

// ECDSA

func encode(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return string(pemEncoded), string(pemEncodedPub)
}

func decode(pemEncoded string, pemEncodedPub string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	block, _ := pem.Decode([]byte(pemEncoded))
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)

	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(*ecdsa.PublicKey)

	return privateKey, publicKey
}

func Read_ECDSA_keys() (string, string) {
	b, err := ioutil.ReadFile(PublicKeyPath) // just pass the file name
	Check_erorr(err)

	pemEncodedPub := string(b) // convert content to a 'string'

	b1, err := ioutil.ReadFile(PrivateKeyPath)
	Check_erorr(err)

	pemEncodedPriv := string(b1) // convert content to a 'string'

	return pemEncodedPriv, pemEncodedPub

}
func Load_ECDSA_keys() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	pemEncodedPriv, pemEncodedPub := Read_ECDSA_keys()

	privateKey, publicKey := decode(pemEncodedPriv, pemEncodedPub)
	return privateKey, publicKey
}

func Write_ECDSA_keys(Priv string, Pub string) {
	//save encPub
	f, err := os.Create(PublicKeyPath)
	Check_erorr(err)
	defer f.Close()
	n1, err := f.WriteString(Pub)
	_ = n1 //avoiding compilation error
	Check_erorr(err)
	f.Sync()

	//save encPriv
	f2, err := os.Create(PrivateKeyPath)
	Check_erorr(err)
	defer f2.Close()
	n2, err := f2.WriteString(Priv)
	_ = n2 //avoiding compilation error
	Check_erorr(err)
	f2.Sync()
}

func Save_keypair_ECDSA() {

	if !FileExists(PublicKeyPath) {

		privateKey, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
		publicKey := &privateKey.PublicKey

		encPriv, encPub := encode(privateKey, publicKey)
		Write_ECDSA_keys(encPriv, encPub)

	}
}
