package dto

type CategoryPageQueryDTO struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Name     string `json:"name"`
	Type     int    `json:"type"`
}
