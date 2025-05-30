package service

import (
	"sky-take-out/common/result"
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/vo"
)

type OrderService interface {
	SubmitOrder(dto dto.OrdersSubmitDTO) (vo.OrderSubmitVO, error)
	PageQueryForUser(page int, pageSize int, status int) (result.PageResult, error)
	OrderDetail(id int) (vo.OrderVO, error)
	Payment(outTradeNo string) error
	ConditionSearch(queryDTO dto.OrdersPageQueryDTO) (result.PageResult, error)
	Statistics() (vo.OrderStatisticsVO, error)
	Details(id int) (vo.OrderVO, error)
	Confirm(cancelDTO dto.OrdersCancelDTO) error
	Rejection(rejectionDTO dto.OrdersRejectionDTO) error
	Cancel(cancelDTO dto.OrdersCancelDTO) error
	Delivery(id int) error
	Complete(id int) error
	Reminder(id int) error
	Repetition(id int) error
}
