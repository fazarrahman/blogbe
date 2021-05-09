package main

import (
	"log"

	"github.com/fazarrahman/blogbe/auth"
	dbcfg "github.com/fazarrahman/blogbe/config/mongodb"
	auth_rest "github.com/fazarrahman/blogbe/delivery/authRest"
	delivery_rest "github.com/fazarrahman/blogbe/delivery/rest"
	user_db_repo "github.com/fazarrahman/blogbe/domain/user/repository/mongodb"
	"github.com/fazarrahman/blogbe/lib"
	"github.com/fazarrahman/blogbe/service"

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
	g.Use(corsInit())
	delivery_rest.New(svc).Register(g.Group("/api"))
	auth_rest.New(svc).Register(g.Group("/api/auth"))

	g.Run(lib.GetEnv("APP_PORT"))
}

func envInit() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
	}
}

func corsInit() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
