package routes

import (
	"github.com/gin-gonic/gin"
	"sky-take-out/resources/controllerParams"
	"sky-take-out/server/interceptor"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(interceptor.JwtTokenAdminInterceptor())
	r.POST("/admin/employee/login", controllerParams.AdminEmployeeController.Login)

	return r
}
