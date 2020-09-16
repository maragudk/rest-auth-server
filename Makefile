.PHONY: certs demo start

certs: generate_cert.go
	go run generate_cert.go --rsa-bits=2048 --host=localhost

demo:
	http -v --session=./session.json --verify=no https://localhost:8080/check
	http -v --session=./session.json --form --verify=no post https://localhost:8080/signup name=demo password=1234567890
	http -v --session=./session.json --form --verify=no post https://localhost:8080/login name=demo password=1234567890
	http -v --session=./session.json --verify=no https://localhost:8080/check
	http -v --session=./session.json --form --verify=no post https://localhost:8080/logout
	http -v --session=./session.json --verify=no https://localhost:8080/check

generate_cert.go:
	wget "https://raw.githubusercontent.com/golang/go/master/src/crypto/tls/generate_cert.go"

start: certs
	go run cmd/ras/*.go
