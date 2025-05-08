package vo

import (
	"time"
)

type OrderSubmitVO struct {
	ID          int64     `json:"id"`
	OrderNumber string    `json:"orderNumber"`
	OrderAmount float64   `json:"orderAmount"`
	OrderTime   time.Time `json:"orderTime"`
}
