package main

import (
	"log"

	"github.com/romanlryji/todo-app"
	"github.com/romanlryji/todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while running http server %s\n", err.Error())
	}
}
