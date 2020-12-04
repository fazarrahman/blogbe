package authRest

import (
	"blogbe/auth"
	d "blogbe/delivery"
	"blogbe/helper"
	"blogbe/model"
	"blogbe/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Rest ...
type AuthRest struct {
	Svc *service.Svc
}

// New ...
func New(_svc *service.Svc) *AuthRest {
	return &AuthRest{Svc: _svc}
}

// Register ...
func (r *AuthRest) Register(g *gin.RouterGroup) {
	g.POST("/login", r.Login)
	g.GET("/tokeninfo", d.MustAuthorize(), r.GetTokenInfo)
}

// GetTokenInfo ..
func (r *AuthRest) GetTokenInfo(c *gin.Context) {
	ti, exists := auth.GetTokenInfo(c)
	if exists {
		c.JSON(http.StatusOK, ti)
		return
	}
	log.Fatalln("error")
}

// Login ..
func (r *AuthRest) Login(c *gin.Context) {
	var req model.User
	c.BindJSON(&req)
	res, err := r.Svc.CheckUsernamePassword(c, &service.UserPasswordCheckRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if res == nil && err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Println(err)
	if *res == false {
		c.JSON(http.StatusBadRequest, "Invalid password")
		return
	}

	c.Request.ParseForm()
	c.Request.Form.Add("client_id", helper.GetEnv("AUTH_CLIENT_ID"))
	c.Request.Form.Add("client_secret", helper.GetEnv("AUTH_SECRET"))
	c.Request.Form.Add("scope", "read")
	c.Request.Form.Add("grant_type", "password")
	c.Request.Form.Add("username", req.Username)
	c.Request.Form.Add("password", req.Password)

	auth.GetAccessToken(c)
}
