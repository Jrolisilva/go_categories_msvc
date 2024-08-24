package main

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := r.Group("/categories")

	categoryRoutes.POST("/", controllers.CreateCategory)
}