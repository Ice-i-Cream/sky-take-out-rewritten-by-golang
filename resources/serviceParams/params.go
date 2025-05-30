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
var ShoppingCartService service.ShoppingCartService = new(impl.ShoppingCartServiceImpl)
var AddressBookService service.AddressBookService = new(impl.AddressBookServiceImpl)
var OrderService service.OrderService = new(impl.OrderServiceImpl)
var ReportService service.ReportService = new(impl.ReportServiceImpl)
var WorkSpaveService service.WorkSpaceService = new(impl.WorkSpaceServiceImpl)
