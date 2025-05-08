package mapper

import (
	"log"
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/commonParams"
	"strings"
)

type OrderDetailMapper struct{}

func (o *OrderDetailMapper) InsertBatch(list []entity.OrderDetail) error {
	insertSQL := "insert into order_detail (name, image, order_id, dish_id, setmeal_id, dish_flavor, amount) values "
	args := []interface{}{}

	for _, order := range list {
		insertSQL = insertSQL + "(?, ?, ?, ?, ?, ?, ?),"
		args = append(args,
			order.Name, order.Image, order.OrderID, order.DishID,
			order.SetmealID, order.DishFlavor, order.Amount,
		)
	}
	insertSQL = strings.TrimSuffix(insertSQL, ",")
	log.Println(insertSQL, args)
	_, err := commonParams.Tx.Exec(insertSQL, args...)
	return err
}
