package serviceParams

import (
	"sky-take-out/server/service"
	"sky-take-out/server/service/impl"
)

var EmployeeService service.EmployeeService = new(impl.EmployeeServiceImpl)
var CategoryService service.CategoryService = new(impl.CategoryServiceImpl)
var DishService service.DishService = new(impl.DishServiceImpl)
var SetmealService service.SetmealService = new(impl.SetmealServiceImpl)
var UserService service.UserService = new(impl.UserServiceImpl)
