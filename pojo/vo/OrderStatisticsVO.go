package vo

type OrderStatisticsVO struct {
	ToBeConfirmed      int64 `json:"toBeConfirmed"`
	Confirmed          int64 `json:"confirmed"`
	DeliveryInProgress int64 `json:"deliveryInProgress"`
}
