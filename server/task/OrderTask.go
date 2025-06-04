package task

import (
	"log"
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/mapperParams"
	"time"
)

type OrderTask struct{}

func (o *OrderTask) ProcessTimeoutOrder() {
	log.Println("定时处理超时订单：%s", time.Now().Format("2006-01-02 15:04:05"))
	t := time.Now().Add(time.Minute * -15)
	ordersList, err := mapperParams.OrderMapper.GetByStatusAndOrderTimeLT(entity.PENDING_PAYMENT, t)
	if err != nil {
		log.Println(err)
		return
	}
	if len(ordersList) > 0 {
		for _, order := range ordersList {
			order.Status = entity.CANCELLED
			order.CancelReason = "订单超时，自动取消"
			order.CancelTime = time.Now()
			err := mapperParams.OrderMapper.Update(order)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
	return
}

func (o *OrderTask) ProcessDeliveryOrder() {
	log.Println("定时处理派送中的订单")
	t := time.Now().Add(time.Minute * -60)
	ordersList, err := mapperParams.OrderMapper.GetByStatusAndOrderTimeLT(entity.DELIVERY_IN_PROGRESS, t)
	if err != nil {
		log.Println(err)
		return
	}
	if len(ordersList) > 0 {
		for _, order := range ordersList {
			order.Status = entity.COMPLETED
			err := mapperParams.OrderMapper.Update(order)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
	return
}
