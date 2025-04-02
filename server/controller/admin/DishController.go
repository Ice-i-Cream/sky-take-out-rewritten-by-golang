package admin

import (
	"github.com/gin-gonic/gin"
	"log"
	"sky-take-out/pojo/dto"
	"sky-take-out/resources/functionParams"
	"sky-take-out/resources/serviceParams"
	"strings"
)

type DishController struct{}

func (d *DishController) Save(ctx *gin.Context) {
	log.Println("新增菜品")
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		var dishDTO dto.DishDTO
		err = ctx.ShouldBindJSON(&dishDTO)
		if err != nil {
			return nil, err
		}
		return nil, serviceParams.DishService.SaveWithFlavor(dishDTO)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (d *DishController) Page(ctx *gin.Context) {
	log.Println("菜品分页查询")
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		dishPageQueryDTO := dto.DishPageQueryDTO{
			Status:     functionParams.ToInt(ctx.Query("status")),
			CategoryID: functionParams.ToInt(ctx.Query("categoryId")),
			Page:       functionParams.ToInt(ctx.Query("page")),
			PageSize:   functionParams.ToInt(ctx.Query("pageSize")),
			Name:       ctx.Query("name"),
		}

		return serviceParams.DishService.PageQuery(dishPageQueryDTO)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)

}

func (d *DishController) Delete(ctx *gin.Context) {
	log.Println("菜品批量删除")
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		list := strings.Split(ctx.Query("ids"), ",")
		return nil, serviceParams.DishService.DeleteBatch(list)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (d *DishController) FindById(ctx *gin.Context) {
	log.Println("根据id查询菜品")
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		id := functionParams.ToInt(ctx.Param("id"))
		return serviceParams.DishService.GetByIdWithFlavor(id)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)

}

func (d *DishController) FindByIds(ctx *gin.Context) {

}
