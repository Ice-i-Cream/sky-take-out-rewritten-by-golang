package dto

import (
	"time"
)

type OrdersPageQueryDTO struct {
	Page      int       `json:"page"`
	PageSize  int       `json:"pageSize"`
	Number    string    `json:"number"`
	Phone     string    `json:"phone"`
	Status    int       `json:"status"` // Changed from *int to int
	BeginTime time.Time `json:"beginTime" time_format:"2006-01-02 15:04:05"`
	EndTime   time.Time `json:"endTime" time_format:"2006-01-02 15:04:05"`
	UserId    int64     `json:"userId"` // Changed from *int64 to int64
}
