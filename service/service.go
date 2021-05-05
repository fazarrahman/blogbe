package service

import (
	userRepository "blogbe/domain/user/repository"
	"blogbe/error"
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
	GetAccessToken(c *gin.Context) *error.Error
	GetUserByID(c *gin.Context, id string) (*model.User, *error.Error)
	InsertUser(c *gin.Context, r *model.User) *error.Error
	CheckUsernamePassword(ctx *gin.Context, r *UserPasswordCheckRequest) (*bool, *error.Error)
}
