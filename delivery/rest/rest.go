package rest

import (
	d "blogbe/delivery"
	"blogbe/error"
	"blogbe/service"

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

}

// GetUser ...
func (r *Rest) GetUser(c *gin.Context) {
	qry := c.Request.URL.Query()

	if qry.Get("username") == "" {
		c.JSON(http.StatusBadRequest, error.BadRequest("Username is required"))
		return
	}

	u, err := r.Svc.GetUser(c, qry.Get("username"))
	if err != nil {
		c.JSON(err.StatusCode, err)
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
