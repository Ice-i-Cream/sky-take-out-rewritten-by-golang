package vo

type OrderOverViewVO struct {
	WaitingOrders int64 `json:"waitingOrders"`

	DeliveredOrders int64 `json:"deliveredOrders"`

	CompletedOrders int64 `json:"completedOrders"`

	CancelledOrders int64 `json:"cancelledOrders"`

	AllOrders int64 `json:"allOrders"`
}
