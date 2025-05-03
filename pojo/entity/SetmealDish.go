package entity

type SetmealDish struct {
	ID        int64   `json:"id"`        // ID of the relationship
	SetmealID int64   `json:"setmealId"` // ID of the set meal
	DishID    int64   `json:"dishId"`    // ID of the dish
	Name      string  `json:"name"`      // Name of the dish (redundant field)
	Price     float64 `json:"price"`     // Original price of the dish
	Copies    int     `json:"copies"`    // Number of copies
}
