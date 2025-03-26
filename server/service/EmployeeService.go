package service

import (
	"sky-take-out/common/result"
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/vo"
)

type EmployeeService interface {
	Login(dto dto.EmployeeLoginDTO) (vo.EmployeeLoginVO, error)
	Save(dto dto.EmployeeDTO) error
	PageQuery(queryDTO dto.EmployeePageQueryDTO) (result.PageResult, error)
}
