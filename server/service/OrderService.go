package service

import (
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/vo"
)

type OrderService interface {
	SubmitOrder(dto dto.OrdersSubmitDTO) (vo.OrderSubmitVO, error)
}
