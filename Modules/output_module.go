package Modules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//output structure
type Response struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
	Pubkey    string `json:"pubkey"`
}

//makes the output prettier
func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

//generate output
func Generate_output(message string) []byte {

	//calling Check_input_length
	isValidInputLength := Check_input_length(message)

	//checking the returned verdict
	if !isValidInputLength {
		panic("Input erorr, The message must be 250 characters or less")
	}

	importedRSAPrivateKey := *LoadRSAPrivatePemKey(PrivateKeyPath)

	signature := SignPKCS1v15(message, importedRSAPrivateKey)

	pubkey, err := ioutil.ReadFile(PublicKeyPath)
	Check_erorr(err)

	var response = Response{message, signature, string(pubkey)}

	var jsonData []byte
	jsonData, err = json.Marshal(response)
	Check_erorr(err)

	jsonData, err = prettyprint(jsonData)
	Check_erorr(err)
	return jsonData
}

//generate and print output to command line
func Print_output(message string) {

	jsonData := Generate_output(message)

	fmt.Println(string(jsonData))

}

// ECDSA

//generate output
func Generate_output_ECDSA(message string) []byte {

	//calling Check_input_length
	isValidInputLength := Check_input_length(message)

	//checking the returned verdict
	if !isValidInputLength {
		panic("Input erorr, The message must be 250 characters or less")
	}

	privateKey, _ := Load_ECDSA_keys()

	signature := SignECDSA(message, privateKey)

	b, err := ioutil.ReadFile(PublicKeyPath) // just pass the file name
	Check_erorr(err)

	publicKey := string(b)
	var response = Response{message, signature, publicKey}

	var jsonData []byte
	jsonData, err = json.Marshal(response)
	Check_erorr(err)

	jsonData, err = prettyprint(jsonData)
	Check_erorr(err)
	return jsonData
}

func Print_output_ECDSA(message string) {
	jsonData := Generate_output_ECDSA(message)

	fmt.Println(string(jsonData))
}
