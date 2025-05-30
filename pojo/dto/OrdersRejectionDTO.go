package dto

type OrdersRejectionDTO struct {
	Id              int    `json:"id"`
	RejectionReason string `json:"rejectionReason"`
}
