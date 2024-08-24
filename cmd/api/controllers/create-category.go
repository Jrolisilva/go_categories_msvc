package controllers

import (
	"net/http"
	use_cases "github.com/Jrolisilva/gocatgories/internal/use-cases"

	"github.com/gin-gonic/gin"
)

type createCategoryInout {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(ctx *gin.Context) {
	var body createCategoryInout

	if err := ctx.ShouldbindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"success": false, "message": err.Error()})
		return
	}
	useCase := use_cases.NewCreateCategoryUseCase()
	err := useCase.Execute(body.Name)

	if err != nil {
		ctx.JSON(400, gin.H{"success": false, "message": err.Error()})
		return
	}
	ctx.JSON(201, gin.H{"success": true, "message": "Category created"})
}