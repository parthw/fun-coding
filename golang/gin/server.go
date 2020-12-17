package main

import (
	"example.com/gin/controllers"
	"example.com/gin/service"
	"github.com/gin-gonic/gin"
)

var (
	postService    service.PostService        = service.New()
	postController controllers.PostController = controllers.New(postService)
)

func main() {
	server := gin.Default()

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	server.GET("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, postController.FindAll())
	})

	server.POST("/save", func(ctx *gin.Context) {
		ctx.JSON(200, postController.Save(ctx))
	})

	server.Run(":8080")
}
