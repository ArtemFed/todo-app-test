package main

import (
	"github.com/ArtemFed/todo-app-test"
	"github.com/ArtemFed/todo-app-test/pkg/handler"
	"github.com/ArtemFed/todo-app-test/pkg/repository"
	"github.com/ArtemFed/todo-app-test/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	myHandler := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run(viper.GetString("8000"), myHandler.InitRoutes()); err != nil {
		log.Fatalf("error occered while running server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
