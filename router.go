package main

import (
	"go_gin/controller"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {

	r.POST("/api/auth/register", controller.Register)
	return r
}
