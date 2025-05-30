package admin

import (
	"github.com/gin-gonic/gin"
	"sky-take-out/pojo/dto"
	"sky-take-out/resources/functionParams"
	"sky-take-out/resources/serviceParams"
	"time"
)

type OrderController struct{}

func (o *OrderController) ConditionSearch(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		beginTime, _ := time.Parse("2006-01-02 15:04:05", ctx.Query("beginTime"))
		endTime, _ := time.Parse("2006-01-02 15:04:05", ctx.Query("endTime"))

		ordersPageQueryDTO := dto.OrdersPageQueryDTO{
			Page:      functionParams.ToInt(ctx.Query("page")),
			PageSize:  functionParams.ToInt(ctx.Query("pageSize")),
			Status:    functionParams.ToInt(ctx.Query("status")),
			Number:    ctx.Query("number"),
			Phone:     ctx.Query("phone"),
			BeginTime: beginTime,
			EndTime:   endTime,
			UserId:    -1,
		}

		return serviceParams.OrderService.ConditionSearch(ordersPageQueryDTO)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (o *OrderController) Statistics(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		orderStatisticsVO, err := serviceParams.OrderService.Statistics()
		return orderStatisticsVO, err
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (o *OrderController) Details(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		id := functionParams.ToInt(ctx.Param("id"))
		return serviceParams.OrderService.Details(id)

	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)

}

func (o *OrderController) Confirm(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		ordersCancelDTO := dto.OrdersCancelDTO{}
		err = ctx.ShouldBind(&ordersCancelDTO)
		if err != nil {
			return nil, err
		}

		return nil, serviceParams.OrderService.Confirm(ordersCancelDTO)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (o *OrderController) Rejection(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		ordersRejectionDTO := dto.OrdersRejectionDTO{}
		err = ctx.ShouldBind(&ordersRejectionDTO)
		return nil, serviceParams.OrderService.Rejection(ordersRejectionDTO)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (o *OrderController) Cancel(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		ordersCancelDTO := dto.OrdersCancelDTO{}
		err = ctx.ShouldBind(&ordersCancelDTO)
		if err != nil {
			return nil, err
		}
		return nil, serviceParams.OrderService.Cancel(ordersCancelDTO)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (o *OrderController) Delivery(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		id := functionParams.ToInt(ctx.Param("id"))
		return nil, serviceParams.OrderService.Delivery(id)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (o *OrderController) Complete(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		id := functionParams.ToInt(ctx.Param("id"))
		return nil, serviceParams.OrderService.Complete(id)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}
