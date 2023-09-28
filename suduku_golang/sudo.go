package main

import (
	"sudu/router"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	router.RegisterRouter(engine)
	engine.Run(":8082")
}
