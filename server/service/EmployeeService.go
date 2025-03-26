package service

import (
	"sky-take-out/common/result"
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/entity"
	"sky-take-out/pojo/vo"
)

type EmployeeService interface {
	Login(dto dto.EmployeeLoginDTO) (vo.EmployeeLoginVO, error)
	Save(dto dto.EmployeeDTO) error
	PageQuery(dto dto.EmployeePageQueryDTO) (result.PageResult, error)
	StartOrStop(employee entity.Employee) error
	GetById(id int) (entity.Employee, error)
	Update(employee entity.Employee) error
}
