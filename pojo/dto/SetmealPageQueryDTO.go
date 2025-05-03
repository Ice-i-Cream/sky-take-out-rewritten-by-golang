package dto

type SetmealPageQueryDTO struct {
	Page       int    `json:"page"`
	PageSize   int    `json:"pageSize"`
	Name       string `json:"name"`
	CategoryId int    `json:"categoryId"`
	Status     int    `json:"status"`
}
