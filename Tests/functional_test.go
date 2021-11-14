package Tests

import (
	"bufio"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
	"testing"

	"github.com/ahmedsharyo/crypto-sign-challenge/Modules"
)

//tests Check_input_length function
func Test_Check_input_length(t *testing.T) {

	//test case 1
	isValidInputLength := Modules.Check_input_length("Welcome to the Jungle")

	//checking the returned verdict
	if !isValidInputLength {
		t.Errorf("on test case %d", 1) // to indicate test failed

	}

	//test case 2
	input := "Welcome to the Jungl"
	for i := 0; i < 100; i++ {
		input += "eeeeee"
	}
	isValidInputLength = Modules.Check_input_length(input)

	//checking the returned verdict
	if isValidInputLength {
		t.Errorf("on test case %d", 2) // to indicate test failed
	}

}

//-----------------------------------------------------

//tests SignPKCS1v15 function and verify signature
func Test_Signature(t *testing.T) {

	Modules.Save_keypair()

	//test case 1
	message := "Welcome to the Jungl"
	importedRSAPrivateKey := *Modules.LoadRSAPrivatePemKey(Modules.PrivateKeyPath)
	importedRSAPublicKey := LoadPublicPemKey(Modules.PublicKeyPath)

	signature := Modules.SignPKCS1v15(message, importedRSAPrivateKey)
	if signature == "Error from signing" {
		t.Errorf("on test case %d", 1) // to indicate test failed

	}
	verif := VerifyPKCS1v15(signature, message, *importedRSAPublicKey)
	if verif == "Error from verification:" {
		t.Errorf("on test case %d", 1) // to indicate test failed

	}

}

//Load public key stored in PEM format and create RSA public Key using go x509.ParsePKCS1PublicKey
func LoadPublicPemKey(fileName string) *rsa.PublicKey {

	publicKeyFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}

	pemfileinfo, _ := publicKeyFile.Stat()

	size := pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)
	Modules.Check_erorr(err)
	data, _ := pem.Decode([]byte(pembytes))
	publicKeyFile.Close()
	publicKeyFileImported, err := x509.ParsePKCS1PublicKey(data.Bytes)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
	return publicKeyFileImported
}

//verifies an RSA PKCS#1 v1.5 signature. hashed is the result of hashing the input message.
func VerifyPKCS1v15(signature string, plaintext string, pubkey rsa.PublicKey) string {
	sig, _ := base64.StdEncoding.DecodeString(signature)
	hashed := sha256.Sum256([]byte(plaintext))
	err := rsa.VerifyPKCS1v15(&pubkey, crypto.SHA256, hashed[:], sig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from verification: %s\n", err)
		return "Error from verification:"
	}
	return "Signature Verification Passed"
}
