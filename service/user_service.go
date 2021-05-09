package service

import (
	"log"
	"time"

	ue "github.com/fazarrahman/blogbe/domain/user/entity"
	"github.com/fazarrahman/blogbe/error"
	"github.com/fazarrahman/blogbe/model"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
	"golang.org/x/crypto/bcrypt"
	validate "gopkg.in/go-playground/validator.v9"
)

// GetUserByID ...
func (s *Svc) GetUserByID(c *gin.Context, id string) (*model.User, *error.Error) {
	userEntity, err := s.UserRepository.GetUser(c, "_id", id)
	if userEntity == nil && err == nil {
		return nil, error.NotFound("User not found")
	} else if err != nil {
		return nil, error.InternalServerError(err.Error())
	}

	var res model.User
	deepcopier.Copy(userEntity).To(&res)
	return &res, nil
}

/*func (s *Svc) GetUsers(ctx *gin.Context, filter bson.M) ([]*ue.User, error) {
	users, err := s.UserRepository.GetUsers(ctx, filter)
	if err != nil {
		return nil, err
	}
	return users, nil
}*/

// InsertUser ..
func (s *Svc) InsertUser(ctx *gin.Context, r *model.User) *error.Error {
	errs := validate.New().Struct(r)
	if errs != nil {
		log.Println(errs)
		return error.BadRequest(errs.Error())
	}

	pwdhash, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return error.InternalServerError(err.Error())
	}

	// validate username
	us, err := s.UserRepository.GetUser(ctx, "username", r.Username)
	if err != nil {
		log.Println(err)
		return error.InternalServerError(err.Error())
	}
	if us != nil {
		return error.ResourceAlreadyExist("Username already exists")
	}

	// validate email
	eEmail, err := s.UserRepository.GetUser(ctx, "email", r.Email)
	if err != nil {
		log.Println(err)
		return error.InternalServerError(err.Error())
	}
	if eEmail != nil {
		return error.ResourceAlreadyExist("Email is already registered")
	}

	var u ue.User
	deepcopier.Copy(r).To(&u)
	u.Password = pwdhash
	u.Status = 1
	u.CreatedDate = time.Now()
	err = s.UserRepository.InsertUser(ctx, &u)

	if err != nil {
		log.Println(err)
		return error.InternalServerError(err.Error())
	}

	return nil
}
