package vo

import (
	"encoding/json"
	"sky-take-out/pojo/entity"
	"time"
)

// 订单状态 1待付款 2待接单 3已接单 4派送中 5已完成 6已取消
const (
	PENDING_PAYMENT      = 1
	TO_BE_CONFIRMED      = 2
	CONFIRMED            = 3
	DELIVERY_IN_PROGRESS = 4
	COMPLETED            = 5
	CANCELLED            = 6
)

// 支付状态 0未支付 1已支付 2退款
const (
	UN_PAID = 0
	PAID    = 1
	REFUND  = 2
)

// Orders represents an order entity
type OrderVO struct {
	ID                    int64                `json:"id"`
	Number                string               `json:"number"`
	Status                int64                `json:"status"`
	UserID                int64                `json:"userId"`
	AddressBookID         int64                `json:"addressBookId"`
	OrderTime             time.Time            `json:"orderTime"`
	CheckoutTime          time.Time            `json:"checkoutTime"`
	PayMethod             int64                `json:"payMethod"`
	PayStatus             int64                `json:"payStatus"`
	Amount                float64              `json:"amount"`
	Remark                string               `json:"remark"`
	UserName              string               `json:"userName"`
	Phone                 string               `json:"phone"`
	Address               string               `json:"address"`
	Consignee             string               `json:"consignee"`
	CancelReason          string               `json:"cancelReason"`
	RejectionReason       string               `json:"rejectionReason"`
	CancelTime            time.Time            `json:"cancelTime"`
	EstimatedDeliveryTime time.Time            `json:"estimatedDeliveryTime"`
	DeliveryStatus        int64                `json:"deliveryStatus"`
	DeliveryTime          time.Time            `json:"deliveryTime"`
	PackAmount            int64                `json:"packAmount"`
	TablewareNumber       int64                `json:"tablewareNumber"`
	TablewareStatus       int64                `json:"tablewareStatus"`
	OrderDishes           string               `json:"orderDishes"`
	OrderDetailList       []entity.OrderDetail `json:"orderDetailList"`
}

// MarshalJSON implements custom JSON marshaling for time fields
func (o OrderVO) MarshalJSON() ([]byte, error) {
	type Alias OrderVO
	return json.Marshal(&struct {
		OrderTime             string `json:"orderTime"`
		CheckoutTime          string `json:"checkoutTime"`
		CancelTime            string `json:"cancelTime"`
		EstimatedDeliveryTime string `json:"estimatedDeliveryTime"`
		DeliveryTime          string `json:"deliveryTime"`
		Alias
	}{
		OrderTime:             o.OrderTime.Format("2006-01-02 15:04:05"),
		CheckoutTime:          o.CheckoutTime.Format("2006-01-02 15:04:05"),
		CancelTime:            o.CancelTime.Format("2006-01-02 15:04:05"),
		EstimatedDeliveryTime: o.EstimatedDeliveryTime.Format("2006-01-02 15:04:05"),
		DeliveryTime:          o.DeliveryTime.Format("2006-01-02 15:04:05"),
		Alias:                 (Alias)(o),
	})
}

// UnmarshalJSON implements custom JSON unmarshaling for time fields
func (o *OrderVO) UnmarshalJSON(data []byte) error {
	type Alias OrderVO
	aux := &struct {
		OrderTime             string `json:"orderTime"`
		CheckoutTime          string `json:"checkoutTime"`
		CancelTime            string `json:"cancelTime"`
		EstimatedDeliveryTime string `json:"estimatedDeliveryTime"`
		DeliveryTime          string `json:"deliveryTime"`
		*Alias
	}{
		Alias: (*Alias)(o),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	parseTime := func(timeStr string) (time.Time, error) {
		if timeStr == "" {
			return time.Time{}, nil
		}
		return time.Parse("2006-01-02 15:04:05", timeStr)
	}

	var err error
	o.OrderTime, err = parseTime(aux.OrderTime)
	if err != nil {
		return err
	}
	o.CheckoutTime, err = parseTime(aux.CheckoutTime)
	if err != nil {
		return err
	}
	o.CancelTime, err = parseTime(aux.CancelTime)
	if err != nil {
		return err
	}
	o.EstimatedDeliveryTime, err = parseTime(aux.EstimatedDeliveryTime)
	if err != nil {
		return err
	}
	o.DeliveryTime, err = parseTime(aux.DeliveryTime)
	if err != nil {
		return err
	}

	return nil
}
