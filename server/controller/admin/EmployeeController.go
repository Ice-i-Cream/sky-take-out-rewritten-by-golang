package admin

import (
	"github.com/gin-gonic/gin"
	"log"
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/commonParams"
	"sky-take-out/resources/functionParams"
	"sky-take-out/resources/serviceParams"
	"strconv"
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

func (e *EmployeeController) StartOrStop(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		status, err := strconv.Atoi(ctx.Param("status"))
		if err != nil {
			return nil, err
		}
		id, err := strconv.Atoi(ctx.Query("id"))
		if err != nil {
			return nil, err
		}
		log.Printf("启用禁用员工账号 status = %d id = %d\n", status, id)
		employee := entity.Employee{
			Status: status,
			ID:     int64(id),
			UpdateUser: func() int64 {
				user, _ := commonParams.Thread.Get()["empId"].(float64)
				return int64(user)
			}(),
		}
		err = serviceParams.EmployeeService.StartOrStop(employee)
		return nil, err
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (e *EmployeeController) GetById(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return nil, err
		}
		return serviceParams.EmployeeService.GetById(id)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (e *EmployeeController) Update(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (data interface{}, err error) {
		dto := dto.EmployeeDTO{}
		err = ctx.ShouldBindJSON(&dto)
		if err != nil {
			return nil, err
		}
		employee := entity.Employee{
			ID:         dto.ID,
			Username:   dto.Username,
			Name:       dto.Name,
			Phone:      dto.Phone,
			Sex:        dto.Sex,
			IDNumber:   dto.IDNumber,
			Status:     -1,
			UpdateTime: time.Now(),
			UpdateUser: func() int64 {
				user, _ := commonParams.Thread.Get()["empId"].(float64)
				return int64(user)
			}(),
		}

		return nil, serviceParams.EmployeeService.Update(employee)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}
