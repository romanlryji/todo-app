package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/romanlryji/todo-app"
)

type AuthorizationRepository interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {
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
	}
}
