package auth

import (
	"blogbe/helper"

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
	clientStore.Set(helper.GetEnv("AUTH_CLIENT_ID"), &models.Client{
		ID:     helper.GetEnv("AUTH_CLIENT_ID"),
		Secret: helper.GetEnv("AUTH_SECRET"),
		Domain: helper.GetEnv("APP_DOMAIN"),
	})
	manager.MapClientStorage(clientStore)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	// Initialize the oauth2 service
	ginserver.InitServer(manager)
	ginserver.SetAllowGetAccessRequest(true)
	ginserver.SetClientInfoHandler(server.ClientFormHandler)
}
