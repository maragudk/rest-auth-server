package main

import (
	"log"
	"os"

	"github.com/maragudk/rest-auth-server/server"
	"github.com/maragudk/rest-auth-server/storage"
)

func main() {
	logger := log.New(os.Stdout, "", 0)

	s := server.New(server.Options{
		Address: ":8080",
		Logger:  logger,
		Storer:  storage.New(),
	})

	if err := s.Start(); err != nil {
		logger.Fatalln("Could not start:", err)
	}
}
