package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"sky-take-out/pojo/dto"
	"sky-take-out/resources/functionParams"
)

type UserController struct{}

func (u *UserController) Login(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		var userLoginDTO dto.UserLoginDTO
		err = ctx.ShouldBindJSON(&userLoginDTO)
		if err != nil {
			return nil, err
		}

		log.Println(fmt.Sprintf("微信用户登录：%s", userLoginDTO.Code))
		return nil, nil
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}
