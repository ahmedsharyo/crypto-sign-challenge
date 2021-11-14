run:
	sudo go run main.go -rsa "hello"

test: 
	sudo go test -v ./Test...
	