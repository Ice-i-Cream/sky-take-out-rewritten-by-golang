package admin

import (
	"github.com/gin-gonic/gin"
	"sky-take-out/pojo/dto"
	"sky-take-out/resources/functionParams"
	"sky-take-out/resources/serviceParams"
	"strconv"
)

type CategoryController struct{}

func (c *CategoryController) AddCategory(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		categoryDTO := dto.CategoryDTO{}
		err = ctx.ShouldBind(&categoryDTO)
		if err != nil {
			return nil, err
		}
		return nil, serviceParams.CategoryService.Save(categoryDTO)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (c *CategoryController) Page(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		categoryPageQueryDTO := dto.CategoryPageQueryDTO{
			Name:     ctx.Query("name"),
			PageSize: functionParams.ToInt(ctx.Query("pageSize")),
			Page:     functionParams.ToInt(ctx.Query("page")),
			Type:     functionParams.ToInt(ctx.Query("type")),
		}
		return serviceParams.CategoryService.PageQuery(categoryPageQueryDTO)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (c *CategoryController) DeleteById(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		var id int
		id, err = strconv.Atoi(ctx.Query("id"))
		if err != nil {
			return nil, err
		}
		return nil, serviceParams.CategoryService.DeleteById(id)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (c *CategoryController) Update(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		categoryDTO := dto.CategoryDTO{}
		err = ctx.ShouldBind(&categoryDTO)
		if err != nil {
			return nil, err
		}
		return nil, serviceParams.CategoryService.Update(categoryDTO)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (c *CategoryController) StartOrStop(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		status := functionParams.ToInt(ctx.Param("status"))
		id := functionParams.ToInt(ctx.Query("id"))
		return nil, serviceParams.CategoryService.StartOrStop(status, id)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (c *CategoryController) List(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		return serviceParams.CategoryService.List(functionParams.ToInt(ctx.Query("type")))
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}
