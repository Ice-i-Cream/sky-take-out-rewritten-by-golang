package vo

import (
	"sky-take-out/pojo/entity"
	"time"
)

type SetmealVO struct {
	ID            int64                `json:"id"`
	CategoryID    int64                `json:"categoryId"`
	Name          string               `json:"name"`
	Price         float64              `json:"price"`
	Status        int                  `json:"status"`
	Description   string               `json:"description"`
	Image         string               `json:"image"`
	UpdateTime    time.Time            `json:"-"`
	CategoryName  string               `json:"categoryName"`
	SetmealDishes []entity.SetmealDish `json:"setmealDishes"`
	UTime         string               `json:"updateTime"`
}
