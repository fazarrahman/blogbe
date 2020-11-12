package auth

import (
	"log"

	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

func Init() {
	manager := manage.NewDefaultManager()

	// token store
	manager.MustTokenStorage(store.NewFileTokenStore("data.db"))

	// client store
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	// Initialize the oauth2 service
	ginserver.InitServer(manager)
	ginserver.SetAllowGetAccessRequest(true)
	ginserver.SetClientInfoHandler(server.ClientFormHandler)

	log.Println("Oauth has been initialized")
}

func GetAccessToken(c *gin.Context) {
	ginserver.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		return "theUserId", nil
	})
	ginserver.HandleTokenRequest(c)
}

func GetTokenInfo(c *gin.Context) (interface{}, bool) {
	return c.Get(ginserver.DefaultConfig.TokenKey)

}
