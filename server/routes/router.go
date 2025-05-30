package routes

import (
	"github.com/gin-gonic/gin"
	"sky-take-out/resources/commonParams"
	"sky-take-out/resources/controllerParams"
	"sky-take-out/server/interceptor"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(interceptor.JwtTokenAdminInterceptor())
	r.Use(interceptor.JwtTokenUserInterceptor())
	r.GET("/ws/:sid", func(c *gin.Context) { commonParams.WSServer.ServeWebSocket(c) })
	r.Static("/image", "file/")

	r.POST("/admin/common/upload", controllerParams.AdminCommonController.Upload)

	r.GET("/admin/report/turnoverStatistics", controllerParams.AdminReportController.TurnoverStatistics)
	r.GET("/admin/report/top10", controllerParams.AdminReportController.Top10)
	r.GET("/admin/report/userStatistics", controllerParams.AdminReportController.UserStatistics)
	r.GET("/admin/report/ordersStatistics", controllerParams.AdminReportController.OrderStatistics)
	r.GET("/admin/report/export", controllerParams.AdminReportController.Export)

	r.GET("/admin/workspace/businessData", controllerParams.AdminWorkSpaceController.BusinessData)
	r.GET("/admin/workspace/overviewSetmeals", controllerParams.AdminWorkSpaceController.OverviewSetmeals)
	r.GET("/admin/workspace/overviewDishes", controllerParams.AdminWorkSpaceController.OverviewDishes)
	r.GET("/admin/workspace/overviewOrders", controllerParams.AdminWorkSpaceController.OverviewOrders)

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
	r.DELETE("/admin/dish", controllerParams.AdminDishController.Delete)
	r.GET("/admin/dish/:id", controllerParams.AdminDishController.FindById)
	r.PUT("/admin/dish", controllerParams.AdminDishController.Update)
	r.GET("/admin/dish/list", controllerParams.AdminDishController.List)
	r.POST("/admin/dish/status/:status", controllerParams.AdminDishController.StartOrStop)

	r.GET("/admin/setmeal/page", controllerParams.AdminSetmealController.Page)
	r.POST("/admin/setmeal", controllerParams.AdminSetmealController.Save)
	r.DELETE("/admin/setmeal", controllerParams.AdminSetmealController.Delete)
	r.GET("/admin/setmeal/:id", controllerParams.AdminSetmealController.GetById)
	r.POST("/admin/setmeal/status/:status", controllerParams.AdminSetmealController.StartOrStop)
	r.PUT("/admin/setmeal", controllerParams.AdminSetmealController.Update)

	r.PUT("/admin/shop/:status", controllerParams.AdminShopController.SetStatus)
	r.GET("/admin/shop/status", controllerParams.AdminShopController.GetStatus)

	r.GET("/admin/order/conditionSearch", controllerParams.AdminOrderController.ConditionSearch)
	r.GET("/admin/order/statistics", controllerParams.AdminOrderController.Statistics)
	r.GET("/admin/order/details/:id", controllerParams.AdminOrderController.Details)
	r.PUT("/admin/order/confirm", controllerParams.AdminOrderController.Confirm)
	r.PUT("/admin/order/rejection", controllerParams.AdminOrderController.Rejection)
	r.PUT("/admin/order/cancel", controllerParams.AdminOrderController.Cancel)
	r.PUT("/admin/order/delivery/:id", controllerParams.AdminOrderController.Delivery)
	r.PUT("admin/order/complete/:id", controllerParams.AdminOrderController.Complete)

	r.GET("/user/dish/list", controllerParams.UserDishController.List)

	r.GET("/user/category/list", controllerParams.UserCategoryController.List)

	r.GET("/user/shop/status", controllerParams.UserShopController.GetStatus)

	r.POST("/user/user/login", controllerParams.UserUserController.Login)

	r.GET("/user/setmeal/dish/:id", controllerParams.UserSetmealController.DishList)
	r.GET("/user/setmeal/list", controllerParams.UserSetmealController.List)

	r.GET("/user/shoppingCart/list", controllerParams.UserShoppingCartController.List)
	r.POST("/user/shoppingCart/add", controllerParams.UserShoppingCartController.Add)
	r.POST("/user/shoppingCart/sub", controllerParams.UserShoppingCartController.Sub)
	r.DELETE("/user/shoppingCart/clean", controllerParams.UserShoppingCartController.Clean)

	r.GET("/user/addressBook/list", controllerParams.UserAddressBookController.List)
	r.GET("/user/addressBook/default", controllerParams.UserAddressBookController.GetDefault)
	r.POST("/user/addressBook", controllerParams.UserAddressBookController.Save)
	r.GET("/user/addressBook/:id", controllerParams.UserAddressBookController.GetById)
	r.PUT("/user/addressBook", controllerParams.UserAddressBookController.Update)
	r.PUT("/user/addressBook/default", controllerParams.UserAddressBookController.SetDefault)
	r.DELETE("/user/addressBook", controllerParams.UserAddressBookController.Delete)

	r.POST("/user/order/submit", controllerParams.UserOrderController.Submit)
	r.GET("/user/order/historyOrders", controllerParams.UserOrderController.Page)
	r.GET("/user/order/orderDetail/:id", controllerParams.UserOrderController.OrderDetail)
	r.PUT("/user/order/payment", controllerParams.UserOrderController.Payment)
	r.GET("/user/order/reminder/:id", controllerParams.UserOrderController.Reminder)
	r.POST("user/order/repetition/:id", controllerParams.UserOrderController.Repetition)

	return r
}
