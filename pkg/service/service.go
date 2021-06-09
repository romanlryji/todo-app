package service

import (
	"github.com/romanlryji/todo-app"
	"github.com/romanlryji/todo-app/pkg/repository"
)

type AuthorizationService interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo.UpdateListRequest) error
}

type TodoItem interface {
	Create(userId, listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateItemRequest) error
}

type Service struct {
	AuthorizationService
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		AuthorizationService: NewAuthService(repos.AuthorizationRepository),
		TodoList:             NewTodoListService(repos.TodoList),
		TodoItem:             NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
