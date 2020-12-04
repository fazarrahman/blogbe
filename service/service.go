package service

import (
	userRepository "blogbe/domain/user/repository"
	"blogbe/model"

	"github.com/gin-gonic/gin"
)

// Svc ...
type Svc struct {
	UserRepository userRepository.Repository
}

// New ...
func New(_userRepo userRepository.Repository) *Svc {
	return &Svc{UserRepository: _userRepo}
}

// Service ...
type Service interface {
	GetUser(c *gin.Context, username string) (*model.User, error)
	InsertUser(c *gin.Context, r *model.User) error
	CheckUsernamePassword(c *gin.Context, r *UserPasswordCheckRequest) (*bool, error)
}
