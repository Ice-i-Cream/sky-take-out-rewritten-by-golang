package mapper

import (
	"fmt"
	"log"
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/commonParams"
	"sky-take-out/resources/functionParams"
	"strings"
	"time"
)

type OrderMapper struct{}

func (o *OrderMapper) Insert(orders entity.Orders) (entity.Orders, error) {
	insertSQL := " insert into orders (number, status, user_id, address_book_id, order_time, checkout_time, pay_method, pay_status, amount, remark, phone, address, user_name, consignee, cancel_reason, rejection_reason, cancel_time, estimated_delivery_time, delivery_status, delivery_time, pack_amount,tableware_number, tableware_status) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	args := []interface{}{
		orders.Number, orders.Status, orders.UserID, orders.AddressBookID,
		orders.OrderTime, orders.CheckoutTime, orders.PayMethod, orders.PayStatus,
		orders.Amount, orders.Remark, orders.Phone, orders.Address, orders.UserName,
		orders.Consignee, orders.CancelReason, orders.RejectionReason,
		orders.CancelTime, orders.EstimatedDeliveryTime, orders.DeliveryStatus,
		orders.DeliveryTime, orders.PackAmount, orders.TablewareNumber,
		orders.TablewareStatus,
	}
	log.Println(insertSQL, args)
	exec, err := functionParams.ExecSQL(insertSQL, args)
	if err != nil {
		return entity.Orders{}, err
	}
	orders.ID, err = exec.LastInsertId()
	return orders, err
}

func (o *OrderMapper) PageQuery(dto dto.OrdersPageQueryDTO) (page []interface{}, err error) {
	selectSQL := "select * from orders where true"
	args := []interface{}{}
	if dto.Number != "" {
		selectSQL = selectSQL + " and number like concat('%',?,'%')"
		args = append(args, dto.Number)
	}
	if dto.Phone != "" {
		selectSQL = selectSQL + " and phone like concat('%',?,'%')"
		args = append(args, dto.Phone)
	}
	if dto.UserId != -1 {
		selectSQL = selectSQL + " and user_id = ?"
		args = append(args, dto.UserId)
	}
	if dto.BeginTime != *new(time.Time) {
		selectSQL = selectSQL + " and order_time >= ? "
		args = append(args, dto.BeginTime)
	}
	if dto.EndTime != *new(time.Time) {
		selectSQL = selectSQL + " and order_time <= ? "
		args = append(args, dto.EndTime)
	}
	if dto.Status != -1 {
		selectSQL = selectSQL + " and status = ? "
		args = append(args, dto.Status)
	}
	selectSQL = selectSQL + " order by order_time desc limit ? offset ?"
	args = append(args, dto.PageSize)
	args = append(args, (dto.Page-1)*dto.PageSize)
	log.Println(selectSQL, args)
	rows, err := commonParams.Db.Query(selectSQL, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		orders := entity.Orders{}
		err := rows.Scan(&orders.ID, &orders.Number, &orders.Status, &orders.UserID, &orders.AddressBookID, &orders.OrderTime, &orders.CheckoutTime, &orders.PayMethod, &orders.PayStatus, &orders.Amount, &orders.Remark, &orders.Phone, &orders.Address, &orders.UserName, &orders.Consignee, &orders.CancelReason, &orders.RejectionReason, &orders.CancelTime, &orders.EstimatedDeliveryTime, &orders.DeliveryStatus, &orders.DeliveryTime, &orders.PackAmount, &orders.TablewareNumber, &orders.TablewareStatus)
		if err != nil {
			return nil, err
		}
		page = append(page, orders)
	}

	return page, err
}

func (o *OrderMapper) GetById(id int64) (entity.Orders, error) {
	selectSQL := "select * from orders where id=?"
	args := []interface{}{id}
	log.Println(selectSQL, args)
	rows, err := commonParams.Db.Query(selectSQL, args...)
	if err != nil {
		return entity.Orders{}, err
	}
	orders := entity.Orders{}
	for rows.Next() {
		err = rows.Scan(&orders.ID, &orders.Number, &orders.Status, &orders.UserID, &orders.AddressBookID, &orders.OrderTime, &orders.CheckoutTime, &orders.PayMethod, &orders.PayStatus, &orders.Amount, &orders.Remark, &orders.Phone, &orders.Address, &orders.UserName, &orders.Consignee, &orders.CancelReason, &orders.RejectionReason, &orders.CancelTime, &orders.EstimatedDeliveryTime, &orders.DeliveryStatus, &orders.DeliveryTime, &orders.PackAmount, &orders.TablewareNumber, &orders.TablewareStatus)
	}
	return orders, err
}

func (o *OrderMapper) GetByNumberAndUserId(no string, id int64) (entity.Orders, error) {
	selectSQL := "select * from orders where number=? and user_id=?"
	args := []interface{}{no, id}
	log.Println(selectSQL, args)
	rows, err := commonParams.Db.Query(selectSQL, args...)
	if err != nil {
		return entity.Orders{}, err
	}
	orders := entity.Orders{}
	for rows.Next() {
		err = rows.Scan(&orders.ID, &orders.Number, &orders.Status, &orders.UserID, &orders.AddressBookID, &orders.OrderTime, &orders.CheckoutTime, &orders.PayMethod, &orders.PayStatus, &orders.Amount, &orders.Remark, &orders.Phone, &orders.Address, &orders.UserName, &orders.Consignee, &orders.CancelReason, &orders.RejectionReason, &orders.CancelTime, &orders.EstimatedDeliveryTime, &orders.DeliveryStatus, &orders.DeliveryTime, &orders.PackAmount, &orders.TablewareNumber, &orders.TablewareStatus)

	}
	return orders, err
}

func (o *OrderMapper) Update(orders entity.Orders) (err error) {
	updateSQL := "update orders set"
	var args []interface{}
	if orders.CancelReason != "" {
		updateSQL = updateSQL + " cancel_reason = ?,"
		args = append(args, orders.CancelReason)
	}
	if orders.RejectionReason != "" {
		updateSQL = updateSQL + " rejection_reason = ?,"
		args = append(args, orders.RejectionReason)
	}
	if orders.CancelTime != *new(time.Time) {
		updateSQL = updateSQL + " cancel_time = ?,"
		args = append(args, orders.CancelTime)
	}
	if orders.PayStatus != -1 {
		updateSQL = updateSQL + " pay_status = ?,"
		args = append(args, orders.PayStatus)
	}
	if orders.PayMethod != -1 {
		updateSQL = updateSQL + " pay_method = ?,"
		args = append(args, orders.PayMethod)
	}
	if orders.CheckoutTime != *new(time.Time) {
		updateSQL = updateSQL + " checkout_time = ?,"
		args = append(args, orders.CheckoutTime)
	}
	if orders.Status != -1 {
		updateSQL = updateSQL + " status = ?,"
		args = append(args, orders.Status)
	}
	if orders.DeliveryTime != *new(time.Time) {
		updateSQL = updateSQL + " delivery_time = ?,"
		args = append(args, orders.DeliveryTime)
	}
	updateSQL = strings.TrimSuffix(updateSQL, ",") + " where id=?"
	args = append(args, orders.ID)
	log.Println(updateSQL, args)
	_, err = commonParams.Db.Exec(updateSQL, args...)
	return err
}

func (o *OrderMapper) CountStatus(toBeConfirmed int) (count int64, err error) {
	selectSQL := "select count(*) from orders where status=?"
	args := []interface{}{toBeConfirmed}
	log.Println(selectSQL, args)
	rows, err := commonParams.Db.Query(selectSQL, args...)
	if err != nil {
		return count, err
	}
	for rows.Next() {
		err = rows.Scan(&count)
	}
	return count, err
}

func (o *OrderMapper) SumByMap(m map[interface{}]interface{}) ([]string, error) {
	begin := m["begin"].(time.Time)
	end := m["end"].(time.Time)
	status := m["status"].(int)

	selectSQL := "SELECT SUM(amount) AS total FROM orders WHERE order_time >= ? AND order_time < ? AND status = ?"
	log.Println(selectSQL, begin, end, status)
	list := []string{}
	for t := begin; !t.After(end); t = t.Add(24 * time.Hour) {
		var total float64
		total = 0
		args := []interface{}{t, t.Add(24 * time.Hour), status}
		rows, err := commonParams.Db.Query(selectSQL, args...)
		if err != nil {
			return []string{}, err
		}
		for rows.Next() {
			err = rows.Scan(&total)
		}

		list = append(list, fmt.Sprintf("%f", total))
	}

	return list, nil
}

func (o *OrderMapper) GetSalesTop(begin time.Time, end time.Time) (goodsList []dto.GoodsSalesDTO, err error) {
	args := []interface{}{}
	selectSQL := "select od.name ,sum(od.number) number from order_detail od ,orders o where od.order_id=o.id and o.status=5"
	if begin != *new(time.Time) {
		selectSQL = selectSQL + " and o.order_time >= ?"
		args = append(args, begin)
	}
	if end != *new(time.Time) {
		selectSQL = selectSQL + " and o.order_time < ?"
		args = append(args, end.Add(24*time.Hour))
	}
	selectSQL = selectSQL + " group by od.name order by number desc limit 0,10"
	log.Println(selectSQL, args)
	rows, err := commonParams.Db.Query(selectSQL, args...)
	if err != nil {
		return goodsList, err
	}
	for rows.Next() {
		goods := dto.GoodsSalesDTO{}
		err := rows.Scan(&goods.Name, &goods.Number)
		if err != nil {
			return goodsList, err
		}
		goodsList = append(goodsList, goods)
	}
	return goodsList, err
}

func (o *OrderMapper) CountByMap(m map[interface{}]interface{}) (int64, error) {
	begin := m["begin"].(time.Time)
	end := m["end"].(time.Time)
	status := m["status"].(int)
	selectSQL := "select count(id) as count from orders where true"
	args := []interface{}{}
	if begin != *new(time.Time) {
		selectSQL = selectSQL + " and order_time >= ?"
		args = append(args, begin)
	}
	if end != *new(time.Time) {
		selectSQL = selectSQL + " and order_time < ?"
		args = append(args, end)
	}
	if status != -1 {
		selectSQL = selectSQL + " and status = ?"
		args = append(args, status)
	}
	//log.Println(selectSQL, args)
	rows, err := commonParams.Db.Query(selectSQL, args...)
	if err != nil {
		return 0, err
	}
	var count int64
	for rows.Next() {
		err = rows.Scan(&count)
	}
	return count, nil
}

func (o *OrderMapper) SumAllByMap(m map[interface{}]interface{}) (float64, error) {
	begin := m["begin"].(time.Time)
	end := m["end"].(time.Time)
	status := m["status"].(int)

	selectSQL := "select sum(amount) as total FROM orders where true"
	args := []interface{}{}
	if begin != *new(time.Time) {
		selectSQL = selectSQL + " and order_time >= ?"
		args = append(args, begin)
	}
	if end != *new(time.Time) {
		selectSQL = selectSQL + " and order_time < ?"
		args = append(args, end)
	}
	if status != -1 {
		selectSQL = selectSQL + " and status = ?"
		args = append(args, status)
	}
	//log.Println(selectSQL, args)
	rows, err := commonParams.Db.Query(selectSQL, args...)
	if err != nil {
		return 0, err
	}
	var total float64
	for rows.Next() {
		err = rows.Scan(&total)
	}
	return total, nil
}

func (o *OrderMapper) GetByStatusAndOrderTimeLT(status int, t time.Time) ([]entity.Orders, error) {
	selectSQL := "select * from orders where status = ? and order_time < ?"
	args := []interface{}{status, t}
	log.Println(selectSQL, args)
	list := []entity.Orders{}
	rows, err := commonParams.Db.Query(selectSQL, args...)
	if err != nil {
		return list, err
	}

	for rows.Next() {
		orders := entity.Orders{}
		err = rows.Scan(&orders.ID, &orders.Number, &orders.Status, &orders.UserID, &orders.AddressBookID, &orders.OrderTime, &orders.CheckoutTime, &orders.PayMethod, &orders.PayStatus, &orders.Amount, &orders.Remark, &orders.Phone, &orders.Address, &orders.UserName, &orders.Consignee, &orders.CancelReason, &orders.RejectionReason, &orders.CancelTime, &orders.EstimatedDeliveryTime, &orders.DeliveryStatus, &orders.DeliveryTime, &orders.PackAmount, &orders.TablewareNumber, &orders.TablewareStatus)
		list = append(list, orders)
	}
	return list, err
}
