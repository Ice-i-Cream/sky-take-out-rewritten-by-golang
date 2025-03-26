package admin

import (
	"github.com/gin-gonic/gin"
	"log"
	"sky-take-out/pojo/dto"
	"sky-take-out/resources/functionParams"
	"sky-take-out/resources/serviceParams"
	"strconv"
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

func (e *EmployeeController) Logout(ctx *gin.Context) {
	functionParams.PostProcess(ctx, nil, nil)
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
		employeePageQueryDTO := dto.EmployeePageQueryDTO{
			Name: ctx.Query("name"),
			PageSize: func() int {
				pageSize, err := strconv.Atoi(ctx.Query("pageSize"))
				if err != nil {
					return 0
				}
				return pageSize
			}(),
			Page: func() int {
				page, err := strconv.Atoi(ctx.Query("page"))
				if err != nil {
					return 0
				}
				return page
			}(),
		}

		log.Println("员工分页查询")
		return serviceParams.EmployeeService.PageQuery(employeePageQueryDTO)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}
