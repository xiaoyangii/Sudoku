package router

import (
	"sudu/handler"
	"sudu/handler/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {

	r.Use(middleware.Cors())
	r.Group("/")
	{
		r.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})
		r.GET("sudo", handler.GetSudoku)
	}

}
