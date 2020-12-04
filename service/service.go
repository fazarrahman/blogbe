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
	GetAccessToken(c *gin.Context)
	GetUser(c *gin.Context, username string) (*model.User, error)
	InsertUser(c *gin.Context, r *model.User) error
	CheckUsernamePassword(ctx *gin.Context, r *UserPasswordCheckRequest) (*bool, error)
}
