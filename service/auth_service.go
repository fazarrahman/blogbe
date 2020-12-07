package service

import (
	errorlib "blogbe/error"
	"blogbe/helper"

	"log"

	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"golang.org/x/crypto/bcrypt"
)

// GetAccessTokenRequest ...
type GetAccessTokenRequest struct {
	Username string `validate:"required,min=1"`
	Password string `validate:"required,min=1"`
}

// UserPasswordCheckRequest ...
type UserPasswordCheckRequest struct {
	Username string `validate:"required,min=1"`
	Password string `validate:"required,min=1"`
}

// GetAccessToken ..
func (s *Svc) GetAccessToken(c *gin.Context, req *GetAccessTokenRequest) *errorlib.Error {
	res, err := s.CheckUsernamePassword(c, &UserPasswordCheckRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if res == nil && err != nil {
		return err
	}

	if *res == false {
		return errorlib.BadRequest("Invalid password")
	}

	c.Request.ParseForm()
	c.Request.Form.Add("client_id", helper.GetEnv("AUTH_CLIENT_ID"))
	c.Request.Form.Add("client_secret", helper.GetEnv("AUTH_SECRET"))
	c.Request.Form.Add("scope", "read")
	c.Request.Form.Add("grant_type", "password")
	c.Request.Form.Add("username", req.Username)
	c.Request.Form.Add("password", req.Password)

	ginserver.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		us, err := s.UserRepository.GetUser(c, "username", username)
		if err != nil {
			log.Println(err)
			return "", err
		}

		return us.ID, nil
	})

	ginserver.HandleTokenRequest(c)
	return err
}

// CheckUsernamePassword ..
func (s *Svc) CheckUsernamePassword(ctx *gin.Context, r *UserPasswordCheckRequest) (*bool, *errorlib.Error) {
	userEntity, err := s.UserRepository.GetUser(ctx, "username", r.Username)
	var res bool
	if userEntity == nil && err == nil {
		return nil, errorlib.NotFound("No user data found")
	} else if err != nil {
		return nil, errorlib.InternalServerError(err.Error())
	}

	err = bcrypt.CompareHashAndPassword(userEntity.Password, []byte(r.Password))
	if err != nil {
		res = false
		return &res, errorlib.InternalServerError(err.Error())
	}
	res = true
	return &res, nil
}
