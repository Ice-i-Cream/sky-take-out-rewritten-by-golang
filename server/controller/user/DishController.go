package user

import (
	"github.com/gin-gonic/gin"
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/commonParams"
	"sky-take-out/resources/functionParams"
	"sky-take-out/resources/serviceParams"
)

type DishController struct{}

func (d *DishController) List(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		categoryId := functionParams.ToInt(ctx.Query("categoryId"))
		data = []entity.Dish{}
		commonParams.Do = func() (interface{}, error) { return serviceParams.DishService.List(categoryId) }
		return functionParams.Cache("dishCache", int64(categoryId), data)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}
