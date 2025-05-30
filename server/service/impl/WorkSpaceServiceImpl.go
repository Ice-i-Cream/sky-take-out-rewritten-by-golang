package impl

import (
	"math"
	"sky-take-out/common/constant"
	"sky-take-out/pojo/entity"
	"sky-take-out/pojo/vo"
	"sky-take-out/resources/mapperParams"
	"time"
)

type WorkSpaceServiceImpl struct {
}

func (w *WorkSpaceServiceImpl) GetBusinessData(begin time.Time, end time.Time) (vo.BusinessDataVO, error) {
	m := map[interface{}]interface{}{
		"begin":  begin,
		"end":    end,
		"status": -1,
	}
	totalOrderCount, err := mapperParams.OrderMapper.CountByMap(m)
	if err != nil {
		return vo.BusinessDataVO{}, err
	}
	m["status"] = entity.COMPLETED
	validOrderCount, err := mapperParams.OrderMapper.CountByMap(m)
	if err != nil {
		return vo.BusinessDataVO{}, err
	}

	var orderCompleteRate float64
	if totalOrderCount != 0 {
		orderCompleteRate = float64(validOrderCount) / float64(totalOrderCount)
	} else {
		orderCompleteRate = 0
	}

	totalNumber, err := mapperParams.OrderMapper.SumAllByMap(m)
	if err != nil {
		return vo.BusinessDataVO{}, err
	}

	var unitPrice float64
	if validOrderCount == 0 {
		unitPrice = 0
	} else {
		unitPrice = float64(totalNumber) / float64(validOrderCount)
		unitPrice = math.Round(unitPrice*100) / 100
	}

	newUsers, err := mapperParams.UserMapper.CountByMap(m)
	if err != nil {
		return vo.BusinessDataVO{}, err
	}

	return vo.BusinessDataVO{
		Turnover:            totalNumber,
		ValidOrderCount:     validOrderCount,
		OrderCompletionRate: orderCompleteRate,
		UnitPrice:           unitPrice,
		NewUsers:            newUsers,
	}, nil
}

func (w *WorkSpaceServiceImpl) CountSetmeals() (setmealOverViewVO vo.SetmealOverViewVO, err error) {
	m := map[interface{}]interface{}{
		"categoryId": -1,
	}
	m["status"] = constant.ENABLE
	setmealOverViewVO.Sold, err = mapperParams.SetmealMapper.CountByMap(m)
	m["status"] = constant.DISABLE
	setmealOverViewVO.Discontinued, err = mapperParams.SetmealMapper.CountByMap(m)

	return setmealOverViewVO, err

}

func (w *WorkSpaceServiceImpl) CountDishes() (dishOverViewVO vo.DishOverViewVO, err error) {
	m := map[interface{}]interface{}{
		"categoryId": -1,
	}
	m["status"] = constant.ENABLE
	dishOverViewVO.Sold, err = mapperParams.DishMapper.CountByMap(m)
	m["status"] = constant.DISABLE
	dishOverViewVO.Discontinued, err = mapperParams.DishMapper.CountByMap(m)

	return dishOverViewVO, err
}

func (w *WorkSpaceServiceImpl) CountOrders() (orderOverViewVO vo.OrderOverViewVO, err error) {
	m := map[interface{}]interface{}{
		"begin": time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local),
		"end":   *new(time.Time),
	}
	m["status"] = -1
	orderOverViewVO.AllOrders, err = mapperParams.OrderMapper.CountByMap(m)
	m["status"] = entity.CANCELLED
	orderOverViewVO.CancelledOrders, err = mapperParams.OrderMapper.CountByMap(m)
	m["status"] = entity.COMPLETED
	orderOverViewVO.CompletedOrders, err = mapperParams.OrderMapper.CountByMap(m)
	m["status"] = entity.CONFIRMED
	orderOverViewVO.DeliveredOrders, err = mapperParams.OrderMapper.CountByMap(m)
	m["status"] = entity.TO_BE_CONFIRMED
	orderOverViewVO.WaitingOrders, err = mapperParams.OrderMapper.CountByMap(m)
	return orderOverViewVO, err
}
