package main

import (
	"sudu/dal/cache"
	"sudu/router"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	cache.Init()
	router.RegisterRouter(engine)
	engine.Run(":8082")
}
