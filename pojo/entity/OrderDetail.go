package entity

// OrderDetail represents an order detail entity
type OrderDetail struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	OrderID    int64   `json:"orderId"`
	DishID     int64   `json:"dishId"`
	SetmealID  int64   `json:"setmealId"`
	DishFlavor string  `json:"dishFlavor"`
	Number     int     `json:"number"`
	Amount     float64 `json:"amount"`
	Image      string  `json:"image"`
}
