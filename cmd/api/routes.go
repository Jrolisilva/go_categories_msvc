package main

import (
	"github.com/Jrolisilva/go-catgories-msvc/cmd/api/controllers"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")

	categoryRoutes.POST("/", controllers.CreateCategory)
}