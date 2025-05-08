package dto

type ShoppingCartDTO struct {
	DishId     int64  `json:"dishId"`
	SetmealId  int64  `json:"setmealId"`
	DishFlavor string `json:"dishFlavor"`
}
