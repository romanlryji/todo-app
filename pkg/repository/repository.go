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
}

type TodoItem interface {
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
	}
}
