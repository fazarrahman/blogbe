package main

import (
	"blogbe/auth"
	dbcfg "blogbe/config/mongodb"
	auth_rest "blogbe/delivery/authRest"
	delivery_rest "blogbe/delivery/rest"
	user_db_repo "blogbe/domain/user/repository/mongodb"
	"blogbe/helper"
	"blogbe/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	envInit()

	db, err := dbcfg.New()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Database has been initialized")

	auth.Init()
	log.Println("Oauth has been initialized")

	userDb := user_db_repo.New(db)
	svc := service.New(userDb)

	g := gin.Default()
	delivery_rest.New(svc).Register(g.Group("/api"))
	auth_rest.New(svc).Register(g.Group("/auth"))

	g.Run(helper.GetEnv("APP_PORT"))
}

func envInit() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
	}
}
