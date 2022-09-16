package service

import (
	galEnt "github.com/fazarrahman/blogbe/domain/gallery/entity"
	galleryRepository "github.com/fazarrahman/blogbe/domain/gallery/repository"
	userRepository "github.com/fazarrahman/blogbe/domain/user/repository"
	"github.com/fazarrahman/blogbe/error"
	"github.com/fazarrahman/blogbe/model"

	"github.com/gin-gonic/gin"
)

// Svc ...
type Svc struct {
	UserRepository    userRepository.Repository
	GalleryRepository galleryRepository.Repository
}

// New ...
func New(_userRepo userRepository.Repository, _galleryRepo galleryRepository.Repository) *Svc {
	return &Svc{UserRepository: _userRepo, GalleryRepository: _galleryRepo}
}

// Service ...
type Service interface {
	GetAccessToken(c *gin.Context) *error.Error
	GetUserByID(c *gin.Context, id string) (*model.User, *error.Error)
	InsertUser(c *gin.Context, r *model.User) *error.Error
	CheckUsernamePassword(ctx *gin.Context, r *UserPasswordCheckRequest) (*bool, *error.Error)
	GetGalleries(ctx *gin.Context) ([]*galEnt.Gallery, *error.Error)
}
