package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"sky-take-out/pojo/dto"
	"sky-take-out/resources/functionParams"
	"sky-take-out/resources/serviceParams"
)

type OrderController struct{}

func (o *OrderController) Submit(ctx *gin.Context) {
	log.Println("用户下单")
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		var ordersSubmitDTO dto.OrdersSubmitDTO
		err = ctx.ShouldBind(&ordersSubmitDTO)
		if err != nil {
			return nil, err
		}
		return serviceParams.OrderService.SubmitOrder(ordersSubmitDTO)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (o *OrderController) Payment(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		ordersPaymentDTO := dto.OrdersPaymentDTO{}
		err = ctx.ShouldBind(&ordersPaymentDTO)
		if err != nil {
			return nil, err
		}
		log.Println("模拟用户支付", ordersPaymentDTO)
		return nil, serviceParams.OrderService.Payment(ordersPaymentDTO.OrderNumber)

	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (o *OrderController) Reminder(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		id := functionParams.ToInt(ctx.Param("id"))
		return nil, serviceParams.OrderService.Reminder(id)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (o *OrderController) Page(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		page := functionParams.ToInt(ctx.Query("page"))
		pageSize := functionParams.ToInt(ctx.Query("pageSize"))
		status := functionParams.ToInt(ctx.Query("status"))
		return serviceParams.OrderService.PageQueryForUser(page, pageSize, status)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (o *OrderController) Repetition(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		id := functionParams.ToInt(ctx.Param("id"))
		return nil, serviceParams.OrderService.Repetition(id)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (o *OrderController) OrderDetail(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		orderID := functionParams.ToInt(ctx.Param("id"))
		return serviceParams.OrderService.OrderDetail(orderID)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)

}
