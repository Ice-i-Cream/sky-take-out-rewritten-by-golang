package mapper

import (
	"fmt"
	"log"
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/commonParams"
	"strings"
)

type SetmealDishMapper struct{}

func (m SetmealDishMapper) InsertBatch(dishes []entity.SetmealDish) error {
	if len(dishes) == 0 {
		return fmt.Errorf("insert batch fail")
	}
	var valueStrings []string
	var args []interface{}
	for _, dish := range dishes {
		valueStrings = append(valueStrings, "(?,?,?,?,?)")
		args = append(args, dish.SetmealID, dish.DishID, dish.Name, dish.Price, dish.Copies)
	}

	insertSQL := fmt.Sprintf("INSERT INTO setmeal_dish (setmeal_id, dish_id, name, price, copies) VALUES %s", strings.Join(valueStrings, ","))

	log.Println(insertSQL, args)
	_, err := commonParams.Db.Exec(insertSQL, args...)
	return err
}

func (m SetmealDishMapper) DeleteBySetmealId(id int) error {
	deleteSQL := "delete from setmeal_dish where setmeal_id=?"
	_, err := commonParams.Tx.Exec(deleteSQL, id)
	return err
}

func (m SetmealDishMapper) GetBySetmealId(id int) (list []entity.SetmealDish, err error) {
	selectSQL := "select * from setmeal_dish where setmeal_id=?"
	rows, err := commonParams.Db.Query(selectSQL, id)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		row := entity.SetmealDish{}
		err = rows.Scan(&row.ID, &row.SetmealID, &row.DishID, &row.Name, &row.Price, &row.Copies)
		if err != nil {
			return nil, err
		}
		list = append(list, row)
	}
	return list, err
}
