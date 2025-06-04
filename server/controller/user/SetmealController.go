package user

import (
	"github.com/gin-gonic/gin"
	"sky-take-out/common/constant"
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/commonParams"
	"sky-take-out/resources/functionParams"
	"sky-take-out/resources/serviceParams"
)

type SetmealController struct{}

func (s *SetmealController) List(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		setmeal := entity.Setmeal{
			CategoryId: int64(functionParams.ToInt(ctx.Query("categoryId"))),
			Status:     constant.ENABLE,
		}
		data = []entity.Setmeal{}

		commonParams.Do = func() (interface{}, error) { return serviceParams.SetmealService.List(setmeal) }
		return functionParams.Cache("setmealCache", setmeal.CategoryId, data)

	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (s *SetmealController) DishList(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		id := int64(functionParams.ToInt(ctx.Param("id")))
		return serviceParams.SetmealService.GetDishItemById(id)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}
