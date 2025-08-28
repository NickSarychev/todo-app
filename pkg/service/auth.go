package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/NickSarychev/todo-app"
	"github.com/NickSarychev/todo-app/pkg/repository"
)

const salt = "fgdjklsangk;vnzvxc"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}
func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = genratePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func genratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
