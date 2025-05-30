package vo

type BusinessDataVO struct {
	Turnover            float64 `json:"turnover"`
	ValidOrderCount     int64   `json:"validOrderCount"`
	OrderCompletionRate float64 `json:"orderCompletionRate"`
	UnitPrice           float64 `json:"unitPrice"`
	NewUsers            int64   `json:"newUsers"`
}
