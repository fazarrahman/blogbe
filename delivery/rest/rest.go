package rest

import (
	"errors"
	"log"

	d "github.com/fazarrahman/blogbe/delivery"
	"github.com/fazarrahman/blogbe/lib"
	"github.com/fazarrahman/blogbe/service"

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
	g.GET("/galleries", r.GetGallery)
	g.POST("/galleries", r.AddGallery)

}

// GetUser ...
func (r *Rest) GetUser(c *gin.Context) {
	uid := lib.GetUserIDFromToken(c)
	if uid == nil {
		c.JSON(http.StatusForbidden, "Invalid Token")
	}

	u, err := r.Svc.GetUserByID(c, *uid)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, u)
}

func (r *Rest) GetGallery(c *gin.Context) {
	users, err := r.Svc.GetGalleries(c)
	if err != nil {
		c.JSON(err.StatusCode, err)
	}
	c.JSON(http.StatusOK, users)
}

func (r *Rest) AddGallery(c *gin.Context) {
	type fileName struct {
		Name string `json:"name"`
	}
	var f fileName
	if err := c.BindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("Error parsing body"))
	}
	log.Println(f)
	err := r.Svc.AddGalleries(c, f.Name)
	c.JSON(http.StatusOK, err)
}
