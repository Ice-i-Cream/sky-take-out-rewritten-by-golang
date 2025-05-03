package user

import (
	"github.com/gin-gonic/gin"
	"sky-take-out/resources/functionParams"
	"sky-take-out/resources/serviceParams"
)

type DishController struct{}

func (d *DishController) List(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		categoryId := functionParams.ToInt(ctx.Query("categoryId"))
		return serviceParams.DishService.List(categoryId)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}
