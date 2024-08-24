package controllers

import (
	"net/http"
	use_cases "github.com/Jrolisilva/go-catgories-msvc/internal/use-cases"

	"github.com/gin-gonic/gin"
)

type createCategoryInput {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(ctx *gin.Context) {
	var body createCategoryInput

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