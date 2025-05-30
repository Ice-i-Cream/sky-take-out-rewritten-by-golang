package controllerParams

import (
	"sky-take-out/server/controller/admin"
	"sky-take-out/server/controller/user"
)

var AdminEmployeeController admin.EmployeeController
var AdminCategoryController admin.CategoryController
var AdminDishController admin.DishController
var AdminCommonController admin.CommonController
var AdminSetmealController admin.SetMealController
var AdminShopController admin.ShopController
var AdminOrderController admin.OrderController
var AdminReportController admin.ReportController
var AdminWorkSpaceController admin.WorkSpaceController

var UserDishController user.DishController
var UserShoppingCartController user.ShoppingCartController
var UserShopController user.ShopController
var UserUserController user.UserController
var UserCategoryController user.CategoryController
var UserSetmealController user.SetmealController
var UserAddressBookController user.AddressBookController
var UserOrderController user.OrderController
