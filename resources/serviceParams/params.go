package serviceParams

import (
	"sky-take-out/server/service"
	"sky-take-out/server/service/impl"
)

var EmployeeService service.EmployeeService = new(impl.EmployeeServiceImpl)
