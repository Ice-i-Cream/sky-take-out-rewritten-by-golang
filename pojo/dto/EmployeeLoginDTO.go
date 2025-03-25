package dto

import (
	"github.com/go-playground/validator/v10" // 用于数据验证（可选）
)

type EmployeeLoginDTO struct {
	Username string `json:"username" validate:"required,min=3,max=50"` // 使用validate标签进行数据验证（可选）
	Password string `json:"password" validate:"required,min=6"`        // 使用validate标签进行数据验证（可选）
}

func (e *EmployeeLoginDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(e)
}
