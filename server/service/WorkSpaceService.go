package service

import (
	"sky-take-out/pojo/vo"
	"time"
)

type WorkSpaceService interface {
	GetBusinessData(begin time.Time, end time.Time) (vo.BusinessDataVO, error)
	CountSetmeals() (vo.SetmealOverViewVO, error)
	CountDishes() (vo.DishOverViewVO, error)
	CountOrders() (vo.OrderOverViewVO, error)
}
