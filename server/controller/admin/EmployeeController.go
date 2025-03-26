package admin

import (
	"github.com/gin-gonic/gin"
	"log"
	"sky-take-out/pojo/dto"
	"sky-take-out/resources/functionParams"
	"sky-take-out/resources/serviceParams"
	"sky-take-out/server/service"
)

type EmployeeController struct{}

func (e *EmployeeController) Login(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		var employeeLoginDTO dto.EmployeeLoginDTO
		err = ctx.ShouldBindJSON(&employeeLoginDTO)
		if err != nil {
			return nil, err
		}
		return serviceParams.EmployeeService.Login(employeeLoginDTO)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (e *EmployeeController) Save(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		var employeeDTO dto.EmployeeDTO
		err = ctx.ShouldBindJSON(&employeeDTO)
		if err != nil {
			return nil, err
		}
		log.Println("新增员工：" + employeeDTO.Name)
		err = serviceParams.EmployeeService.Save(employeeDTO)
		return nil, err
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (e *EmployeeController) PageQuery(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		var employeePageQueryDTO dto.EmployeePageQueryDTO
		err = ctx.ShouldBindQuery(&employeePageQueryDTO)
		if err != nil {
			return nil, err
		}
		log.Println("员工分页查询")
		return service.EmployeeService.PageQuery(employeePageQueryDTO)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}
