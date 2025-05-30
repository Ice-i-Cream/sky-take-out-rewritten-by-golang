package dto

type OrdersPaymentDTO struct {
	OrderNumber string `json:"orderNumber"`
	PayMethod   int64  `json:"payMethod"`
}
