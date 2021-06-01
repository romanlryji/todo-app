package main

import (
	"log"

	"github.com/romanlryji/todo-app"
	"github.com/romanlryji/todo-app/pkg/handler"
	"github.com/romanlryji/todo-app/pkg/repository"
	"github.com/romanlryji/todo-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error while reading config %s\n", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while running http server %s\n", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

// docker run --name=todo-db -e POSTGRES_USER=qwerty -e POSTGRES_PASSWORD=qwerty -p 5436:5432 -d --rm postgres
