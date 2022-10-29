package main

import (
	"fmt"
	"github.com/ArtemFed/todo-app-test"
	"github.com/ArtemFed/todo-app-test/pkg/handler"
	"github.com/ArtemFed/todo-app-test/pkg/repository"
	"github.com/ArtemFed/todo-app-test/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	fmt.Println("Step 1")

	if err != nil {
		log.Fatalf("failed to initialize database: %s", err.Error())
	}
	fmt.Println("Step 2")

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	myHandler := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run(viper.GetString("8000"), myHandler.InitRoutes()); err != nil {
		log.Fatalf("error occered while running server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
