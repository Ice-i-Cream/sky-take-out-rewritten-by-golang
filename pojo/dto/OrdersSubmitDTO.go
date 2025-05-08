package dto

import (
	"encoding/json"
	"time"
)

type OrdersSubmitDTO struct {
	AddressBookId         int64     `json:"addressBookId"`
	PayMethod             int64     `json:"payMethod"`
	Remark                string    `json:"remark"`
	EstimatedDeliveryTime time.Time `json:"estimatedDeliveryTime"`
	DeliveryStatus        int64     `json:"deliveryStatus"`
	TablewareNumber       int64     `json:"tablewareNumber"`
	TablewareStatus       int64     `json:"tablewareStatus"`
	PackAmount            int64     `json:"packAmount"`
	Amount                float64   `json:"amount"`
}

// MarshalJSON implements custom JSON marshaling for EstimatedDeliveryTime
func (o OrdersSubmitDTO) MarshalJSON() ([]byte, error) {
	type Alias OrdersSubmitDTO
	return json.Marshal(&struct {
		EstimatedDeliveryTime string `json:"estimatedDeliveryTime"`
		Alias
	}{
		EstimatedDeliveryTime: o.EstimatedDeliveryTime.Format("2006-01-02 15:04:05"),
		Alias:                 (Alias)(o),
	})
}

// UnmarshalJSON implements custom JSON unmarshaling for EstimatedDeliveryTime
func (o *OrdersSubmitDTO) UnmarshalJSON(data []byte) error {
	type Alias OrdersSubmitDTO
	aux := &struct {
		EstimatedDeliveryTime string `json:"estimatedDeliveryTime"`
		*Alias
	}{
		Alias: (*Alias)(o),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	parsedTime, err := time.Parse("2006-01-02 15:04:05", aux.EstimatedDeliveryTime)
	if err != nil {
		return err
	}
	o.EstimatedDeliveryTime = parsedTime
	return nil
}
