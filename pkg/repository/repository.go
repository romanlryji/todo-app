package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/romanlryji/todo-app"
)

type AuthorizationRepository interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo.UpdateListRequest) error
}

type TodoItem interface {
	Create(listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, request todo.UpdateItemRequest) error
}

type Repository struct {
	AuthorizationRepository
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthorizationRepository: NewAuthPostgres(db),
		TodoList:                NewTodoListPostgres(db),
		TodoItem:                NewTodoItemPostgres(db),
	}
}
