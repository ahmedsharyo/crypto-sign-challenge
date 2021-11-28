Crypto Sign
===========

Crypto Sign is an application to sign a message with your private key using RSA/ECDSA algorithms

Installing
----------

This is intended to be used on Unix based file systems.

To install this tool you need to have [Go 1.14][go] installed.

[go]: https://golang.org/


Build
-----
 - Navigate to `/go/src/github.com/<user>/crypto-sign-challenge` directory
 - Run :
        `$ go build`


Usage
-----
 - Navigate to `/go/src/github.com/<user>/crypto-sign-challenge` directory
 - Run :
         `sudo ./crypto-sign-challenge  -algo <MESSAGE>` 

`MESSAGE` is the message you wish to sign with your private key.  The message
must be 250 characters or less. 
  `-algo` is the algorithm to use (RSA or ECDSA )

The following is an example with output included.

```
$ crypto-sign-challenge -rsa 'Welcome to the Jungle'
{
    "message": "Welcome to the Jungle",
    "signature": "MIGIAkIBHEc8FETUYOPze9YxePzBfN2OjbstTYQxfViHu6vziSfDbM5iJ8jCmH3LkScgoTNCRBAMBY407jDC/fYq88iN22cCQgCmytbObfzxtHWHpcYFvOb3PHHDKlv+rtAZJ/+AdxBvihjY/xRDi1PH8GhyEgzW7xzJ1KF7BhqmeMwH9pXUCx6JiA==",
    "pubkey": "-----BEGIN PUBLIC KEY-----\nMIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQAxMXE/k5LOn1ZeSNgILi/fsDyHwwW\nSugmEndN786laNFUJ0Ulzit1FumnY71Op7Gwuqrv+YoqrEwpHtpnV8mLgvEBr9sX\ncNatfZzPtjOLpHzkVfLSCX94E7uNUZx13eigwugCsR87rn94CLRU3GDbLnLO6W4f\n12FkAhynQpvqaWNKpn8=\n-----END PUBLIC KEY-----\n"
}
```

Storage
-------

This project will generate a new private key if it does not exist and will store
it in:

    `$HOME`
    
This project will generate a new public key from the new private key and will store
it in:

    `$HOME` 

Test
-------
- Navigate to `/go/src/github.com/<user>/crypto-sign-challenge` directory
- Run :
       `$ go test -v ./Test...`


Continuous Integration 
-------

How a Continuous Integration system can execute your test(s) Ex. (circle CI)

- Publish the app to a version control paltform Ex. (Github)
- Creat config.yaml file inside (.circleci) directory Ex.

    version: 2
    jobs:
    build:
    docker:
      - image: circleci/golang:1.14
    working_directory: /go/src/github.com/<user>/crypto-sign-challenge
    steps:
      - checkout
      - run: go test -v ./Test...

- visit your https://app.circleci.com/projects 
- add the project and configure the branch name 



Response Examples
---------------------

JSON Schema for your application response:

```json
{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "title": "Signed Identifier",
    "description": "Schema for a signed identifier",
    "type": "object",
    "required": [ "message", "signature", "pubkey" ],
    "properties": {
        "message": {
            "type": "string",
            "description": "original string provided as the input to your app"
        },
        "signature": {
            "type": "string",
            "description": "RFC 4648 compliant Base64 encoded cryptographic signature of the input, calculated using the private key and the SHA256 digest of the input"
        },
        "pubkey": {
            "type": "string",
            "description": "Base64 encoded string (PEM format) of the public key generated from the private key used to create the digital signature"
        }
    }
}
```


EXAMPLE

```
>./your-awesome-app -ecdsa "theAnswerIs42"
```

Returns:

```json
{
    "message":"theAnswerIs42",
    "signature":"MGUCMCDwlFyVdD620p0hRLtABoJTR7UNgwj8g2r0ipNbWPi4Us57YfxtSQJ3dAkHslyBbwIxAKorQmpWl9QdlBUtACcZm4kEXfL37lJ+gZ/hANcTyuiTgmwcEC0FvEXY35u2bKFwhA==",
    "pubkey":"-----BEGIN PUBLIC KEY-----\nMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEI5/0zKsIzou9hL3ZdjkvBeVZFKpDwxTb\nfiDVjHpJdu3+qOuaKYgsLLiO9TFfupMYHLa20IqgbJSIv/wjxANH68aewV1q2Wn6\nvLA3yg2mOTa/OHAZEiEf7bVEbnAov+6D\n-----END PUBLIC KEY-----\n"
}
```


references :
           
           https://8gwifi.org/docs/go-rsa.jsp

           https://pkg.go.dev/crypto/ecdsa
