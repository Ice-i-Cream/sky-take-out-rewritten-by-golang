package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"sky-take-out/pojo/dto"
	"sky-take-out/resources/functionParams"
	"sky-take-out/resources/serviceParams"
)

type ShoppingCartController struct{}

func (s *ShoppingCartController) Add(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		shoppingCartDTO := dto.ShoppingCartDTO{}
		err = ctx.ShouldBind(&shoppingCartDTO)
		if err != nil {
			return nil, err
		}
		log.Printf("添加购物车，商品信息：%d\n", shoppingCartDTO.DishId)
		return nil, serviceParams.ShoppingCartService.AddShoppingCart(shoppingCartDTO)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)

}

func (s *ShoppingCartController) List(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {

		return serviceParams.ShoppingCartService.ShowShoppingCart()
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}
