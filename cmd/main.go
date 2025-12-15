package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "serv", log.LstdFlags|log.Lshortfile)

	serv := server.MyServer(logger)

	if err := serv.Launch(); err != nil {
		logger.Fatal(err)
	}
}
