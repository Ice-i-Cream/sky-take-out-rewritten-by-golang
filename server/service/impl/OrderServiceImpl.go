package impl

import (
	"encoding/json"
	"fmt"
	"log"
	"sky-take-out/common/constant"
	"sky-take-out/common/result"
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

func (o *OrderServiceImpl) OrderDetail(id int) (vo.OrderVO, error) {
	orderDetailList, err := mapperParams.OrderDetailMapper.GetByOrderId(int64(id))
	if err != nil {
		return vo.OrderVO{}, err
	}
	order, err := mapperParams.OrderMapper.GetById(int64(id))
	if err != nil {
		return vo.OrderVO{}, err
	}
	orderVO := vo.OrderVO{
		ID:                    order.ID,
		Number:                order.Number,
		Status:                order.Status,
		UserID:                order.UserID,
		AddressBookID:         order.AddressBookID,
		OrderTime:             order.OrderTime,
		CheckoutTime:          order.CheckoutTime,
		PayMethod:             order.PayMethod,
		PayStatus:             order.PayStatus,
		Amount:                order.Amount,
		Remark:                order.Remark,
		UserName:              order.UserName,
		Phone:                 order.Phone,
		Address:               order.Address,
		Consignee:             order.Consignee,
		CancelReason:          order.CancelReason,
		RejectionReason:       order.RejectionReason,
		CancelTime:            order.CancelTime,
		EstimatedDeliveryTime: order.EstimatedDeliveryTime,
		DeliveryStatus:        order.DeliveryStatus,
		DeliveryTime:          order.DeliveryTime,
		PackAmount:            order.PackAmount,
		TablewareNumber:       order.TablewareNumber,
		TablewareStatus:       order.TablewareStatus,
		OrderDishes:           "",
		OrderDetailList:       orderDetailList,
	}

	return orderVO, nil
}

func (o *OrderServiceImpl) SubmitOrder(dto dto.OrdersSubmitDTO) (orderVO vo.OrderSubmitVO, err error) {
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

func (o *OrderServiceImpl) PageQueryForUser(pageNum int, pageSize int, status int) (result.PageResult, error) {
	ordersPageQueryDTO := dto.OrdersPageQueryDTO{
		Page:     pageNum,
		PageSize: pageSize,
		Status:   status,
		UserId:   int64(commonParams.Thread.Get()["userId"].(float64)),
	}
	page, err := mapperParams.OrderMapper.PageQuery(ordersPageQueryDTO)
	if err != nil {
		return result.PageResult{}, err
	}
	list := []interface{}{}
	for _, orders := range page {
		order := orders.(entity.Orders)
		orderDetailList, err := mapperParams.OrderDetailMapper.GetByOrderId(order.ID)
		if err != nil {
			return result.PageResult{}, err
		}

		orderVO := vo.OrderVO{
			ID:                    order.ID,
			Number:                order.Number,
			Status:                order.Status,
			UserID:                order.UserID,
			AddressBookID:         order.AddressBookID,
			OrderTime:             order.OrderTime,
			CheckoutTime:          order.CheckoutTime,
			PayMethod:             order.PayMethod,
			PayStatus:             order.PayStatus,
			Amount:                order.Amount,
			Remark:                order.Remark,
			UserName:              order.UserName,
			Phone:                 order.Phone,
			Address:               order.Address,
			Consignee:             order.Consignee,
			CancelReason:          order.CancelReason,
			RejectionReason:       order.RejectionReason,
			CancelTime:            order.CancelTime,
			EstimatedDeliveryTime: order.EstimatedDeliveryTime,
			DeliveryStatus:        order.DeliveryStatus,
			DeliveryTime:          order.DeliveryTime,
			PackAmount:            order.PackAmount,
			TablewareNumber:       order.TablewareNumber,
			TablewareStatus:       order.TablewareStatus,
			OrderDishes:           "",
			OrderDetailList:       orderDetailList,
		}

		list = append(list, orderVO)
	}

	return result.PageResult{
		Total:   len(page),
		Records: list,
	}, nil
}

func (o *OrderServiceImpl) Payment(outTradeNo string) error {
	userId := int64(commonParams.Thread.Get()["userId"].(float64))
	orderDB, err := mapperParams.OrderMapper.GetByNumberAndUserId(outTradeNo, userId)
	if err != nil {
		return err
	}

	orders := entity.Orders{
		ID:           orderDB.ID,
		Status:       entity.TO_BE_CONFIRMED,
		PayStatus:    entity.PAID,
		PayMethod:    -1,
		CheckoutTime: time.Now(),
	}
	err = mapperParams.OrderMapper.Update(orders)
	if err != nil {
		return err
	}

	mp := map[string]interface{}{}
	mp["type"] = 1
	mp["orderId"] = orderDB.ID
	mp["content"] = "订单号" + outTradeNo
	jsonBytes, err := json.Marshal(mp)
	if err != nil {
		return err
	}

	jsonString := string(jsonBytes)
	commonParams.WSServer.SendToAllClients(jsonString)

	return nil

}

func (o *OrderServiceImpl) ConditionSearch(dto dto.OrdersPageQueryDTO) (result.PageResult, error) {
	page, err := mapperParams.OrderMapper.PageQuery(dto)
	if err != nil {
		return result.PageResult{}, err
	}
	records, err := getOrderVoList(page)
	return result.PageResult{
		Total:   len(page),
		Records: records,
	}, err

}

func getOrderVoList(page []interface{}) ([]interface{}, error) {
	list := []interface{}{}
	for _, orders := range page {
		order := orders.(entity.Orders)
		orderVO := vo.OrderVO{
			ID:                    order.ID,
			Number:                order.Number,
			Status:                order.Status,
			UserID:                order.UserID,
			AddressBookID:         order.AddressBookID,
			OrderTime:             order.OrderTime,
			CheckoutTime:          order.CheckoutTime,
			PayMethod:             order.PayMethod,
			PayStatus:             order.PayStatus,
			Amount:                order.Amount,
			Remark:                order.Remark,
			UserName:              order.UserName,
			Phone:                 order.Phone,
			Address:               order.Address,
			Consignee:             order.Consignee,
			CancelReason:          order.CancelReason,
			RejectionReason:       order.RejectionReason,
			CancelTime:            order.CancelTime,
			EstimatedDeliveryTime: order.EstimatedDeliveryTime,
			DeliveryStatus:        order.DeliveryStatus,
			DeliveryTime:          order.DeliveryTime,
			PackAmount:            order.PackAmount,
			TablewareNumber:       order.TablewareNumber,
			TablewareStatus:       order.TablewareStatus,
			OrderDishes:           "",
			OrderDetailList:       []entity.OrderDetail{},
		}
		list = append(list, orderVO)
	}
	return list, nil
}

func (o *OrderServiceImpl) Statistics() (vo.OrderStatisticsVO, error) {
	toBeConfirmed, err := mapperParams.OrderMapper.CountStatus(entity.TO_BE_CONFIRMED)
	if err != nil {
		return vo.OrderStatisticsVO{}, err
	}
	confirmed, err := mapperParams.OrderMapper.CountStatus(entity.CONFIRMED)
	if err != nil {
		return vo.OrderStatisticsVO{}, err
	}
	deliveryInProgress, err := mapperParams.OrderMapper.CountStatus(entity.DELIVERY_IN_PROGRESS)
	if err != nil {
		return vo.OrderStatisticsVO{}, err
	}
	orderStatisticsVO := vo.OrderStatisticsVO{
		ToBeConfirmed:      toBeConfirmed,
		Confirmed:          confirmed,
		DeliveryInProgress: deliveryInProgress,
	}
	return orderStatisticsVO, nil
}

func (o *OrderServiceImpl) Details(id int) (vo.OrderVO, error) {
	order, err := mapperParams.OrderMapper.GetById(int64(id))
	if err != nil {
		return vo.OrderVO{}, err
	}
	orderDetailList, err := mapperParams.OrderDetailMapper.GetByOrderId(int64(id))
	if err != nil {
		return vo.OrderVO{}, err
	}
	orderVO := vo.OrderVO{
		ID:                    order.ID,
		Number:                order.Number,
		Status:                order.Status,
		UserID:                order.UserID,
		AddressBookID:         order.AddressBookID,
		OrderTime:             order.OrderTime,
		CheckoutTime:          order.CheckoutTime,
		PayMethod:             order.PayMethod,
		PayStatus:             order.PayStatus,
		Amount:                order.Amount,
		Remark:                order.Remark,
		UserName:              order.UserName,
		Phone:                 order.Phone,
		Address:               order.Address,
		Consignee:             order.Consignee,
		CancelReason:          order.CancelReason,
		RejectionReason:       order.RejectionReason,
		CancelTime:            order.CancelTime,
		EstimatedDeliveryTime: order.EstimatedDeliveryTime,
		DeliveryStatus:        order.DeliveryStatus,
		DeliveryTime:          order.DeliveryTime,
		PackAmount:            order.PackAmount,
		TablewareNumber:       order.TablewareNumber,
		TablewareStatus:       order.TablewareStatus,
		OrderDishes:           "",
		OrderDetailList:       orderDetailList,
	}
	return orderVO, nil

}

func (o *OrderServiceImpl) Confirm(cancelDTO dto.OrdersCancelDTO) error {
	orders := entity.Orders{
		ID:        cancelDTO.Id,
		Status:    entity.CONFIRMED,
		PayStatus: -1,
		PayMethod: -1,
	}
	return mapperParams.OrderMapper.Update(orders)

}

func (o *OrderServiceImpl) Rejection(rejectionDTO dto.OrdersRejectionDTO) error {
	ordersDB, err := mapperParams.OrderMapper.GetById(int64(rejectionDTO.Id))
	if err != nil {
		return err
	}
	if ordersDB == *new(entity.Orders) || !(ordersDB.Status == entity.TO_BE_CONFIRMED) {
		return fmt.Errorf(constant.ORDER_STATUS_ERROR)
	}
	payStatus := ordersDB.PayStatus
	if payStatus == entity.PAID {
		log.Println("申请退款")
	}
	orders := entity.Orders{
		ID:              int64(rejectionDTO.Id),
		Status:          entity.CANCELLED,
		RejectionReason: rejectionDTO.RejectionReason,
		CancelTime:      time.Now(),
		PayMethod:       -1,
		PayStatus:       -1,
	}
	return mapperParams.OrderMapper.Update(orders)

}

func (o *OrderServiceImpl) Cancel(cancelDTO dto.OrdersCancelDTO) error {
	ordersDB, err := mapperParams.OrderMapper.GetById(int64(cancelDTO.Id))
	if err != nil {
		return err
	}
	payStatus := ordersDB.PayStatus
	if payStatus == entity.PAID {
		log.Println("申请退款")
	}
	orders := entity.Orders{
		ID:           cancelDTO.Id,
		Status:       entity.CANCELLED,
		CancelReason: cancelDTO.CancelReason,
		CancelTime:   time.Now(),
		PayStatus:    -1,
		PayMethod:    -1,
	}
	return mapperParams.OrderMapper.Update(orders)
}

func (o *OrderServiceImpl) Delivery(id int) error {
	ordersDB, err := mapperParams.OrderMapper.GetById(int64(id))
	if err != nil {
		return err
	}
	if ordersDB == *new(entity.Orders) || !(ordersDB.Status == entity.CONFIRMED) {
		return fmt.Errorf(constant.ORDER_STATUS_ERROR)
	}
	orders := entity.Orders{
		ID:        ordersDB.ID,
		Status:    entity.DELIVERY_IN_PROGRESS,
		PayStatus: -1,
		PayMethod: -1,
	}
	return mapperParams.OrderMapper.Update(orders)
}

func (o *OrderServiceImpl) Complete(id int) error {
	ordersDB, err := mapperParams.OrderMapper.GetById(int64(id))
	if err != nil {
		return err
	}
	if ordersDB == *new(entity.Orders) || !(ordersDB.Status == entity.DELIVERY_IN_PROGRESS) {
		return fmt.Errorf(constant.ORDER_STATUS_ERROR)
	}
	orders := entity.Orders{
		ID:           ordersDB.ID,
		Status:       entity.COMPLETED,
		DeliveryTime: time.Now(),
		PayStatus:    -1,
		PayMethod:    -1,
	}
	return mapperParams.OrderMapper.Update(orders)
}

func (o *OrderServiceImpl) Reminder(id int) error {
	ordersDB, err := mapperParams.OrderMapper.GetById(int64(id))
	if err != nil {
		return err
	}
	if ordersDB == *new(entity.Orders) {
		return fmt.Errorf(constant.ORDER_STATUS_ERROR)
	}
	mp := map[string]interface{}{}
	mp["type"] = 2
	mp["orderId"] = ordersDB.ID
	mp["content"] = "订单号" + ordersDB.Number
	jsonBytes, err := json.Marshal(mp)
	if err != nil {
		return err
	}

	jsonString := string(jsonBytes)
	commonParams.WSServer.SendToAllClients(jsonString)

	return nil
}

func (o *OrderServiceImpl) Repetition(id int) error {

	userId := int(commonParams.Thread.Get()["userId"].(float64))

	orderDetailList, err := mapperParams.OrderDetailMapper.GetByOrderId(int64(id))
	if err != nil {
		return err
	}
	shoppingCartList := []entity.ShoppingCart{}
	for _, orderDetail := range orderDetailList {
		shoppingCart := entity.ShoppingCart{
			Name:       orderDetail.Name,
			UserID:     int64(userId),
			DishID:     orderDetail.DishID,
			SetmealID:  orderDetail.SetmealID,
			DishFlavor: orderDetail.DishFlavor,
			Number:     orderDetail.Number,
			Amount:     orderDetail.Amount,
			Image:      orderDetail.Image,
			CreateTime: time.Now(),
		}
		shoppingCartList = append(shoppingCartList, shoppingCart)
	}
	commonParams.Tx, err = commonParams.Db.Begin()
	if err != nil {
		return err
	}
	err = mapperParams.ShoppingCartMapper.InsertBatch(shoppingCartList)
	if err != nil {
		commonParams.Tx.Rollback()
		return err
	}
	return commonParams.Tx.Commit()

}
