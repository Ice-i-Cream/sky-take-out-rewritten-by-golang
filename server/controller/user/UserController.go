package user

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"sky-take-out/common/utils"
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/vo"
	"sky-take-out/resources/commonParams"
	"sky-take-out/resources/functionParams"
	"sky-take-out/resources/serviceParams"
	"time"
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

		user, err := serviceParams.UserService.WxLogin(userLoginDTO)
		if err != nil {
			return nil, err
		}

		claims := jwt.MapClaims{
			"claims": jwt.MapClaims{
				"userId": user.ID,
			},
		}
		token, err := utils.GenToken(claims, commonParams.JwtProperties.UserSecretKey)
		if err != nil {
			return nil, err
		}

		commonParams.RedisDb.Set(commonParams.Ctx, token, token, time.Hour)
		userLoginVO := vo.UserLoginVO{
			ID:     user.ID,
			OpenID: user.OpenID,
			Token:  token,
		}
		return userLoginVO, nil
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}
