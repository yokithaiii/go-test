package main

import (
	"log"

	"github.com/yokithaiii/go-test.git"
	"github.com/yokithaiii/go-test.git/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("cannot connect to server: %s", err.Error())
	}
}
