package entity

import (
	"time"
)

type ShoppingCart struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`       // Name
	UserID     int64     `json:"userId"`     // User ID
	DishID     int64     `json:"dishId"`     // Dish ID
	SetmealID  int64     `json:"setmealId"`  // Setmeal ID
	DishFlavor string    `json:"dishFlavor"` // Flavor
	Number     int       `json:"number"`     // Quantity
	Amount     float64   `json:"amount"`     // Amount
	Image      string    `json:"image"`      // Image URL
	CreateTime time.Time `json:"createTime"` // Creation time
}
