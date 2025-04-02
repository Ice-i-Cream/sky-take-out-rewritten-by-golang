package dto

type DishPageQueryDTO struct {
	Page       int    `json:"page"`       // 当前页码
	PageSize   int    `json:"pageSize"`   // 每页显示的记录数
	Name       string `json:"name"`       // 菜品名称
	CategoryID int    `json:"categoryId"` // 分类ID
	Status     int    `json:"status"`     // 状态，0表示禁用，1表示启用，使用指针以表示可选字段
}
