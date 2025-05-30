package mapper

import (
	"log"
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/commonParams"
	"sky-take-out/resources/functionParams"
	"strings"
)

type OrderDetailMapper struct{}

func (o *OrderDetailMapper) InsertBatch(list []entity.OrderDetail) error {
	insertSQL := "insert into order_detail (name, image, order_id, dish_id, setmeal_id, dish_flavor,number, amount) values "
	var args []interface{}

	for _, order := range list {
		insertSQL = insertSQL + "(?, ?, ?, ?, ?, ?, ?, ?),"
		args = append(args,
			order.Name, order.Image, order.OrderID, order.DishID,
			order.SetmealID, order.DishFlavor, order.Number, order.Amount,
		)
	}
	insertSQL = strings.TrimSuffix(insertSQL, ",")
	log.Println(insertSQL, args)
	_, err := functionParams.ExecSQL(insertSQL, args)
	return err
}

func (o *OrderDetailMapper) GetByOrderId(id int64) (list []entity.OrderDetail, err error) {
	sql := "select * from order_detail where order_id = ?"
	args := []interface{}{id}
	log.Println(sql, args)
	rows, err := commonParams.Db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		detail := entity.OrderDetail{}
		err := rows.Scan(&detail.ID, &detail.Name, &detail.Image, &detail.OrderID, &detail.DishID, &detail.SetmealID, &detail.DishFlavor, &detail.Number, &detail.Amount)
		if err != nil {
			return nil, err
		}
		list = append(list, detail)
	}
	return list, err

}
