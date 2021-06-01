package main

import (
	"log"

	"github.com/romanlryji/todo-app"
	"github.com/romanlryji/todo-app/pkg/handler"
	"github.com/romanlryji/todo-app/pkg/repository"
	"github.com/romanlryji/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while running http server %s\n", err.Error())
	}
}
