package rest

import (
	"blogbe/auth"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Rest ...
type Rest struct {
}

// New ...
func New() *Rest {
	return &Rest{}
}

// Register ...
func (r *Rest) Register(g *gin.Engine) {
	auth := g.Group("/oauth2")
	{
		auth.GET("/token", r.GetToken)
		auth.GET("/tokeninfo", MustAuthorize(), r.GetTokenInfo)
	}

	//api := g.Group("/api")
}

// GetToken ..
func (r *Rest) GetToken(c *gin.Context) {
	auth.GetAccessToken(c)
}

// GetTokenInfo ..
func (r *Rest) GetTokenInfo(c *gin.Context) {
	ti, exists := auth.GetTokenInfo(c)
	if exists {
		c.JSON(http.StatusOK, ti)
		return
	}
	log.Fatalln("error")
}
