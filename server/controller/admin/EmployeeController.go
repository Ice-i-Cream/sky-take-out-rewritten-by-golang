package admin

import (
	"github.com/gin-gonic/gin"
	"log"
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/commonParams"
	"sky-take-out/resources/functionParams"
	"sky-take-out/resources/serviceParams"
	"time"
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
			Name:     ctx.Query("name"),
			PageSize: functionParams.ToInt(ctx.Query("pageSize")),
			Page:     functionParams.ToInt(ctx.Query("page")),
		}
		log.Println("员工分页查询")
		return serviceParams.EmployeeService.PageQuery(employeePageQueryDTO)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (e *EmployeeController) StartOrStop(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		employee := entity.Employee{
			Status:     functionParams.ToInt(ctx.Param("status")),
			ID:         int64(functionParams.ToInt(ctx.Query("id"))),
			UpdateUser: functionParams.GetUser(commonParams.Thread.Get()["empId"]),
		}
		log.Printf("启用禁用员工账号 status = %d id = %d\n", employee.Status, employee.ID)
		err = serviceParams.EmployeeService.StartOrStop(employee)
		return nil, err
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (e *EmployeeController) GetById(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		id := functionParams.ToInt(ctx.Param("id"))
		return serviceParams.EmployeeService.GetById(id)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (e *EmployeeController) Update(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		d := dto.EmployeeDTO{}
		err = ctx.ShouldBindJSON(&d)
		if err != nil {
			return nil, err
		}
		employee := entity.Employee{
			ID:         d.ID,
			Username:   d.Username,
			Name:       d.Name,
			Phone:      d.Phone,
			Sex:        d.Sex,
			IDNumber:   d.IDNumber,
			Status:     -1,
			UpdateTime: time.Now(),
			UpdateUser: functionParams.GetUser(commonParams.Thread.Get()["empId"]),
		}

		return nil, serviceParams.EmployeeService.Update(employee)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}
