package routes

import (
	"github.com/gin-gonic/gin"
	"sky-take-out/resources/controllerParams"
	"sky-take-out/server/interceptor"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(interceptor.JwtTokenAdminInterceptor())
	r.Static("/image", "file/")
	r.POST("/admin/employee/login", controllerParams.AdminEmployeeController.Login)
	r.POST("/admin/employee/logout", controllerParams.AdminEmployeeController.Logout)
	r.POST("/admin/employee", controllerParams.AdminEmployeeController.Save)
	r.GET("/admin/employee/page", controllerParams.AdminEmployeeController.PageQuery)
	r.POST("/admin/employee/status/:status", controllerParams.AdminEmployeeController.StartOrStop)
	r.GET("/admin/employee/:id", controllerParams.AdminEmployeeController.GetById)
	r.PUT("/admin/employee", controllerParams.AdminEmployeeController.Update)

	r.GET("/admin/category/page", controllerParams.AdminCategoryController.Page)
	r.POST("/admin/category", controllerParams.AdminCategoryController.AddCategory)
	r.DELETE("/admin/category", controllerParams.AdminCategoryController.DeleteById)
	r.PUT("/admin/category", controllerParams.AdminCategoryController.Update)
	r.POST("/admin/category/status/:status", controllerParams.AdminCategoryController.StartOrStop)
	r.GET("/admin/category/list", controllerParams.AdminCategoryController.List)
	r.GET("/admin/dish/page", controllerParams.AdminDishController.Page)
	r.POST("/admin/dish", controllerParams.AdminDishController.Save)
	r.POST("/admin/common/upload", controllerParams.AdminCommonController.Upload)
	r.DELETE("/admin/dish", controllerParams.AdminDishController.Delete)
	r.GET("/admin/dish/:id", controllerParams.AdminDishController.FindById)

	return r
}
