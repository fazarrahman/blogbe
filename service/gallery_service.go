package service

import (
	galEnt "github.com/fazarrahman/blogbe/domain/gallery/entity"
	"github.com/fazarrahman/blogbe/error"
	"github.com/gin-gonic/gin"
)

func (s *Svc) GetGalleries(ctx *gin.Context) ([]*galEnt.Gallery, *error.Error) {
	galleries, err := s.GalleryRepository.GetGallery(ctx)
	if err != nil {
		return nil, error.InternalServerError(err.Error())
	}
	return galleries, nil
}

func (s *Svc) AddGalleries(ctx *gin.Context, source string) *error.Error {
	if source == "" {
		return error.BadRequest("File name is empty")
	}

	s.GalleryRepository.AddGallery(ctx, source)
	return nil
}
