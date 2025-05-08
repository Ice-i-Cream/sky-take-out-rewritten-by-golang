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
