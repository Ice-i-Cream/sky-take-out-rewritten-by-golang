package dto

type EmployeePageQueryDTO struct {
	Name     string `json:"name"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
}
