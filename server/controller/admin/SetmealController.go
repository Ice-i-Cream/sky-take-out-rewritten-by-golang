package admin

import (
	"github.com/gin-gonic/gin"
	"sky-take-out/pojo/dto"
	"sky-take-out/resources/functionParams"
	"sky-take-out/resources/serviceParams"
	"strings"
)

type SetMealController struct{}

func (s *SetMealController) Page(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		setmealPageQueryDTO := dto.SetmealPageQueryDTO{
			Page:       functionParams.ToInt(ctx.Query("page")),
			PageSize:   functionParams.ToInt(ctx.Query("pageSize")),
			Name:       ctx.Query("name"),
			CategoryId: functionParams.ToInt(ctx.Query("categoryId")),
			Status:     functionParams.ToInt(ctx.Query("status")),
		}
		return serviceParams.SetmealService.PageQuery(setmealPageQueryDTO)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (s *SetMealController) Save(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		var setmealDTO dto.SetmealDTO
		err = ctx.ShouldBind(&setmealDTO)
		if err != nil {
			return nil, err
		}
		return nil, serviceParams.SetmealService.SaveWithDish(setmealDTO)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (s *SetMealController) Delete(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		setmeal := strings.Split(ctx.Query("ids"), ",")
		err = serviceParams.SetmealService.DeleteBatch(setmeal)
		return nil, err
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (s *SetMealController) GetById(ctx *gin.Context) {

	exec := func(ctx *gin.Context) (data interface{}, err error) {
		setmealVO, err := serviceParams.SetmealService.GetByIdWithDish(ctx.Param("id"))
		return setmealVO, err
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (s *SetMealController) Update(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		var setmealDTO dto.SetmealDTO
		err = ctx.ShouldBind(&setmealDTO)
		if err != nil {
			return nil, err
		}
		err = serviceParams.SetmealService.Update(setmealDTO)
		return nil, err
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (s *SetMealController) StartOrStop(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		setmealDTO := dto.SetmealDTO{
			Status: functionParams.ToInt(ctx.Param("status")),
			Id:     int64(functionParams.ToInt(ctx.Query("id"))),
		}

		err = serviceParams.SetmealService.StartOrStop(setmealDTO)
		return nil, err
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}
