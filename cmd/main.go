package main

import (
	"log"

	"github.com/spf13/viper"
	"github.com/yokithaiii/go-test.git"
	"github.com/yokithaiii/go-test.git/pkg/handler"
	"github.com/yokithaiii/go-test.git/pkg/repository"
	"github.com/yokithaiii/go-test.git/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init config file. %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("cannot connect to server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	
	return viper.ReadInConfig()
}
