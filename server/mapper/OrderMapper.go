package mapper

import (
	"log"
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/commonParams"
)

type OrderMapper struct{}

func (m OrderMapper) Insert(orders entity.Orders) (entity.Orders, error) {
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
	exec, err := commonParams.Tx.Exec(insertSQL, args...)
	if err != nil {
		return entity.Orders{}, err
	}
	orders.ID, err = exec.LastInsertId()
	return orders, err
}
