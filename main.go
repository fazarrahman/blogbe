package main

import (
	"blogbe/auth"
	delivery_rest "blogbe/delivery/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	auth.Init()

	g := gin.Default()
	delivery_rest.New().Register(g)

	g.Run(":9096")
}
