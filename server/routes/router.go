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
	r.POST("/admin/employee/logout", controllerParams.AdminEmployeeController.Logout)
	r.POST("/admin/employee", controllerParams.AdminEmployeeController.Save)
	r.GET("/admin/employee/page", controllerParams.AdminEmployeeController.PageQuery)
	r.POST("/admin/employee/status/:status", controllerParams.AdminEmployeeController.StartOrStop)
	r.GET("/admin/employee/:id", controllerParams.AdminEmployeeController.GetById)
	r.PUT("/admin/employee", controllerParams.AdminEmployeeController.Update)

	return r
}
