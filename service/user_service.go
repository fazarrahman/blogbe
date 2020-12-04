package service

import (
	ue "blogbe/domain/user/entity"
	"blogbe/model"
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
	"golang.org/x/crypto/bcrypt"
	validate "gopkg.in/go-playground/validator.v9"
)

// GetUser ...
func (s *Svc) GetUser(c *gin.Context, username string) (*model.User, error) {
	userEntity, err := s.UserRepository.GetUser(c, username)
	if userEntity == nil && err == nil {
		return nil, errors.New("No data")
	} else if err != nil {
		return nil, err
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

func (s *Svc) InsertUser(ctx *gin.Context, r *model.User) error {
	errs := validate.New().Struct(r)
	if errs != nil {
		log.Println(errs)
		return errors.New("Bad Request")
	}

	pwdhash, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// validate usernqme
	us, err := s.UserRepository.GetUser(ctx, r.Username)
	if err != nil {
		return err
	}
	if us != nil {
		return errors.New("Username already exists")
	}

	var u ue.User
	deepcopier.Copy(r).To(&u)
	u.Password = pwdhash
	u.Status = 1
	u.CreatedDate = time.Now()
	err = s.UserRepository.InsertUser(ctx, &u)

	if err != nil {
		return err
	}

	return nil
}
