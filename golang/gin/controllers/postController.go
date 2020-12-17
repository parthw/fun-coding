package controllers

import (
	"example.com/gin/modal"
	"example.com/gin/service"
	"github.com/gin-gonic/gin"
)

//PostController interface
type PostController interface {
	Save(ctx *gin.Context) modal.Post
	FindAll() []modal.Post
}

type controller struct {
	service service.PostService
}

// New controller
func New(service service.PostService) PostController {
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) modal.Post {
	var post modal.Post
	ctx.BindJSON(&post)
	return c.service.Save(post)
}

func (c *controller) FindAll() []modal.Post {
	return c.service.FindAll()
}
