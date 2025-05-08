package impl

import (
	"fmt"
	"sky-take-out/common/constant"
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/entity"
	"sky-take-out/pojo/vo"
	"sky-take-out/resources/commonParams"
	"sky-take-out/resources/functionParams"
	"sky-take-out/resources/mapperParams"
	"strconv"
	"time"
)

type OrderServiceImpl struct{}

func (o OrderServiceImpl) SubmitOrder(dto dto.OrdersSubmitDTO) (orderVO vo.OrderSubmitVO, err error) {
	addressBook, err := mapperParams.AddressBookMapper.GetById(dto.AddressBookId)
	if err != nil {
		return orderVO, fmt.Errorf(constant.ADDRESS_BOOK_IS_NULL)
	}
	userId := functionParams.ToInt(commonParams.Thread.Get()["userId"].(float64))
	shoppingCart := entity.ShoppingCart{
		UserID:    int64(userId),
		SetmealID: -1,
		DishID:    -1,
	}
	shoppingCartList, err := mapperParams.ShoppingCartMapper.List(shoppingCart)
	if err != nil {
		return orderVO, fmt.Errorf(constant.SHOPPING_CART_IS_NULL)
	}

	orders := entity.Orders{
		CancelTime:            time.Now(),
		CheckoutTime:          time.Now(),
		DeliveryTime:          time.Now(),
		AddressBookID:         dto.AddressBookId,
		PayMethod:             dto.PayMethod,
		Remark:                dto.Remark,
		EstimatedDeliveryTime: dto.EstimatedDeliveryTime,
		DeliveryStatus:        dto.DeliveryStatus,
		PackAmount:            dto.PackAmount,
		Amount:                dto.Amount,
		TablewareNumber:       dto.TablewareNumber,
		TablewareStatus:       dto.TablewareStatus,
		OrderTime:             time.Now(),
		PayStatus:             entity.UN_PAID,
		Status:                entity.PENDING_PAYMENT,
		Number:                strconv.FormatInt(time.Now().UnixMilli(), 10),
		Phone:                 addressBook.Phone,
		Consignee:             addressBook.Consignee,
		UserID:                int64(userId),
	}

	commonParams.Tx, err = commonParams.Db.Begin()
	if err != nil {
		return orderVO, err
	}
	orders, err = mapperParams.OrderMapper.Insert(orders)
	if err != nil {
		commonParams.Tx.Rollback()
		return orderVO, err
	}
	var orderDetailList []entity.OrderDetail
	for _, cart := range shoppingCartList {
		orderDetailList = append(orderDetailList, entity.OrderDetail{
			ID:         cart.ID,
			Name:       cart.Name,
			OrderID:    orders.ID,
			DishID:     cart.DishID,
			SetmealID:  cart.SetmealID,
			DishFlavor: cart.DishFlavor,
			Number:     cart.Number,
			Amount:     cart.Amount,
			Image:      cart.Image,
		})
	}

	err = mapperParams.OrderDetailMapper.InsertBatch(orderDetailList)
	if err != nil {
		commonParams.Tx.Rollback()
		return orderVO, err
	}

	err = mapperParams.ShoppingCartMapper.DeleteByUserId(int64(userId))
	if err != nil {
		commonParams.Tx.Rollback()
		return vo.OrderSubmitVO{}, err
	}

	return vo.OrderSubmitVO{
		ID:          orders.ID,
		OrderNumber: orders.Number,
		OrderAmount: orders.Amount,
		OrderTime:   orders.OrderTime,
	}, commonParams.Tx.Commit()
}
