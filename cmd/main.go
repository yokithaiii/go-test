package main

import (
	"log"

	"github.com/yokithaiii/go-test.git"
	"github.com/yokithaiii/go-test.git/pkg/handler"
	"github.com/yokithaiii/go-test.git/pkg/repository"
	"github.com/yokithaiii/go-test.git/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("cannot connect to server: %s", err.Error())
	}
}
