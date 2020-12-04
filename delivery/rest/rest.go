package rest

import (
	d "blogbe/delivery"
	"blogbe/model"
	"blogbe/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Rest ...
type Rest struct {
	Svc *service.Svc
}

// New ...
func New(_svc *service.Svc) *Rest {
	return &Rest{Svc: _svc}
}

// Register ...
func (r *Rest) Register(g *gin.RouterGroup) {
	g.GET("/profile", d.MustAuthorize(), r.GetUser)
	g.POST("/auth/signup", d.MustAuthorize(), r.PostUser)

}

// GetUser ...
func (r *Rest) GetUser(c *gin.Context) {
	qry := c.Request.URL.Query()

	if qry.Get("username") == "" {
		c.JSON(http.StatusBadRequest, errors.New("Username is required"))
		return
	}

	u, err := r.Svc.GetUser(c, qry.Get("username"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, u)
}

/*func (r *Rest) GetUsers(c *gin.Context) {
	users, err := r.Svc.GetUsers(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, users)
}*/

func (r *Rest) PostUser(c *gin.Context) {
	var req model.User
	c.BindJSON(&req)
	err := r.Svc.InsertUser(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "OK")
}

func (r *Rest) CheckUsernamePassword(c *gin.Context) {
	var req service.UserPasswordCheckRequest
	c.BindJSON(&req)
	res, err := r.Svc.CheckUsernamePassword(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, res)
}
