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

var UserDishController user.DishController
var UserShopController user.ShopController
var UserUserController user.UserController
var UserCategoryController user.CategoryController
