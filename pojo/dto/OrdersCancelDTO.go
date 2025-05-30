package dto

type OrdersCancelDTO struct {
	Id           int64  `json:"id"`
	CancelReason string `json:"cancelReason"`
}
