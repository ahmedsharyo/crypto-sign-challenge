package main

import (
	"os"

	"github.com/ahmedsharyo/crypto-sign-challenge/Modules"
)

func RSA(message string) {

	Modules.Save_keypair()

	Modules.Print_output(message)
}

func ECDSA(message string) {

	Modules.Save_keypair_ECDSA()

	Modules.Print_output_ECDSA(message)
}

func main() {

	//takes arguments from command line
	algo := os.Args[1]
	message := os.Args[2]

	if algo == "-rsa" {
		RSA(message)

	} else {

		ECDSA(message)
	}

}
