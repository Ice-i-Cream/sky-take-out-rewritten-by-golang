package interceptor

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sky-take-out/common/utils"
	"sky-take-out/resources/commonParams"
	"sky-take-out/resources/functionParams"
	"strings"
)

func JwtTokenUserInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {

		if !strings.HasPrefix(c.Request.URL.Path, "/user") {
			c.Next()
			return
		}

		excludedPaths := []string{"/", "/user/employee/login", "/user/shop/status"}

		if functionParams.IsExcludedPath([]string{}, excludedPaths, c.Request.URL.Path) {
			tokenString := c.GetHeader(commonParams.JwtProperties.AdminTokenName)

			if commonParams.RedisDb.Get(commonParams.Ctx, tokenString).Val() != tokenString {
				log.Printf("\nError parsing JWT: token失效\n")
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort() // 终止请求
				return
			}

			claims, err := utils.ParseToken(tokenString, commonParams.JwtProperties.AdminSecretKey)
			if err != nil {
				log.Printf("Error parsing JWT: %v\n", err)
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort() // 终止请求
				return
			}
			commonParams.Thread.Set(claims)
			var id = claims["userId"].(float64)
			log.Printf("\njwt校验:%s\n当前用户id:%d\n", tokenString, int(id))
		}

		c.Next()
	}
}
