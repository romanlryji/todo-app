package service

import (
	"github.com/romanlryji/todo-app"
	"github.com/romanlryji/todo-app/pkg/repository"
)

type AuthorizationService interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	AuthorizationService
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		AuthorizationService: NewAuthService(repos.AuthorizationRepository),
	}
}
