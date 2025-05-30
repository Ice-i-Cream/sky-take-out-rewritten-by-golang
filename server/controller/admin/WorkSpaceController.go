package admin

import (
	"github.com/gin-gonic/gin"
	"sky-take-out/resources/functionParams"
	"sky-take-out/resources/serviceParams"
	"time"
)

type WorkSpaceController struct {
}

func (w *WorkSpaceController) BusinessData(ctx *gin.Context) {

	exec := func(ctx *gin.Context) (data interface{}, err error) {
		t := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)
		return serviceParams.WorkSpaveService.GetBusinessData(t, *new(time.Time))
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (w *WorkSpaceController) OverviewSetmeals(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		return serviceParams.WorkSpaveService.CountSetmeals()
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (w *WorkSpaceController) OverviewDishes(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		return serviceParams.WorkSpaveService.CountDishes()
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (w *WorkSpaceController) OverviewOrders(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		return serviceParams.WorkSpaveService.CountOrders()
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}
