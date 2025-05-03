package dto

import "sky-take-out/pojo/entity"

type SetmealDTO struct {
	Id            int64                `json:"id"`
	CategoryId    int64                `json:"categoryId"`
	Name          string               `json:"name"`
	Price         interface{}          `json:"price"`
	Status        int                  `json:"status"`
	Description   string               `json:"description"`
	Image         string               `json:"image"`
	SetmealDishes []entity.SetmealDish `json:"setmealDishes"`
}
