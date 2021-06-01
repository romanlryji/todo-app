package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/romanlryji/todo-app"
	"github.com/romanlryji/todo-app/pkg/repository"
)

const salt = "nc982u9t8ur29nc8t5-28un28uger"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
