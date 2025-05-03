package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"sky-take-out/resources/commonParams"
	"sky-take-out/resources/functionParams"
)

type ShopController struct{}

func (s *ShopController) GetStatus(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		var statusStr string
		status := functionParams.ToInt(commonParams.RedisDb.Get(commonParams.Ctx, "SHOP_STATUS").Val())
		if status == 1 {
			statusStr = "营业中"
		} else {
			statusStr = "打烊中"
		}
		log.Println(fmt.Sprintf("获取店铺的营业状态为：%s", statusStr))
		return status, nil
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}
