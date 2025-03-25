package admin

import (
	"github.com/gin-gonic/gin"
	"sky-take-out/pojo/dto"
	"sky-take-out/resources/functionParams"
	"sky-take-out/resources/serviceParams"
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
