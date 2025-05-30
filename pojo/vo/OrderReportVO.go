package vo

type OrderReportVO struct {
	DateList            string  `json:"dateList"`
	OrderCountList      string  `json:"orderCountList"`
	ValidOrderCountList string  `json:"validOrderCountList"`
	TotalOrderCount     int64   `json:"totalOrderCount"`
	ValidOrderCount     int64   `json:"validOrderCount"`
	OrderCompletionRate float64 `json:"orderCompletionRate"`
}
