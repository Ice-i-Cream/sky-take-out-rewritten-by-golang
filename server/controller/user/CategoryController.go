package user

import (
	"github.com/gin-gonic/gin"
	"sky-take-out/resources/functionParams"
	"sky-take-out/resources/serviceParams"
)

type CategoryController struct{}

func (c *CategoryController) List(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		return serviceParams.CategoryService.List(functionParams.ToInt(ctx.Query("type")))
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}
