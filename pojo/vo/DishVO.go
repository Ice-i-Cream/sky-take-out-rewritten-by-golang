package vo

import (
	"sky-take-out/pojo/entity"
	"time"
)

// DishVO 表示菜品的视图对象
type DishVO struct {
	ID           int64               `json:"id"`           // 菜品ID
	Name         string              `json:"name"`         // 菜品名称
	CategoryID   int64               `json:"categoryId"`   // 菜品分类ID
	Price        float64             `json:"price"`        // 菜品价格，使用float64以表示小数
	Image        string              `json:"image"`        // 图片URL
	Description  string              `json:"description"`  // 描述信息
	Status       int                 `json:"status"`       // 状态，0表示停售，1表示起售
	UpdateTime   time.Time           `json:"-"`            // 更新时间
	CategoryName string              `json:"categoryName"` // 分类名称
	Flavors      []entity.DishFlavor `json:"flavors"`      // 菜品关联的口味
	Time         string              `json:"updateTime"`
	// Copies       int       `json:"copies"`     // 如果需要，可以取消注释
}
