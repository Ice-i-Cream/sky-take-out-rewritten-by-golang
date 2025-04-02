package dto

import "sky-take-out/pojo/entity"

type DishDTO struct {
	ID          int64               `json:"id"`          // 菜品ID
	Name        string              `json:"name"`        // 菜品名称
	CategoryID  int64               `json:"categoryId"`  // 菜品分类ID
	Price       string              `json:"price"`       // 菜品价格，使用float64以表示小数
	Image       string              `json:"image"`       // 图片URL
	Description string              `json:"description"` // 描述信息
	Status      int                 `json:"status"`      // 状态，0表示停售，1表示起售
	Flavors     []entity.DishFlavor `json:"flavors"`     // 菜品关联的口味
}
