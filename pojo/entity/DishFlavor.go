package entity

// 虽然当前结构体未直接使用time，但可以用于记录创建/更新时间（如果需要）

// DishFlavor 表示菜品的口味
type DishFlavor struct {
	ID     int64  `json:"id"`     // 口味ID
	DishID int64  `json:"dishId"` // 菜品ID
	Name   string `json:"name"`   // 口味名称
	Value  string `json:"value"`  // 口味数据，可以是具体描述或编码
}
