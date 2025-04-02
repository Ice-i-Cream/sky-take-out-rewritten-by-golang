package entity

import (
	"time"
)

// Dish 表示菜品实体
type Dish struct {
	ID          int64     `json:"id"`          // 菜品ID
	Name        string    `json:"name"`        // 菜品名称
	CategoryID  int64     `json:"category_id"` // 菜品分类ID
	Price       float64   `json:"price"`       // 菜品价格，使用float64以表示小数
	Image       string    `json:"image"`       // 图片URL
	Description string    `json:"description"` // 描述信息
	Status      int       `json:"status"`      // 状态，0表示停售，1表示起售
	CreateTime  time.Time `json:"create_time"` // 创建时间
	UpdateTime  time.Time `json:"update_time"` // 更新时间
	CreateUser  int64     `json:"create_user"` // 创建用户ID
	UpdateUser  int64     `json:"update_user"` // 更新用户ID
}
