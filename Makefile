.PHONY: certs start

certs: generate_cert.go
	go run generate_cert.go --rsa-bits=2048 --host=localhost

generate_cert.go:
	wget "https://raw.githubusercontent.com/golang/go/master/src/crypto/tls/generate_cert.go"

start: certs
	go run cmd/ras/*.go
