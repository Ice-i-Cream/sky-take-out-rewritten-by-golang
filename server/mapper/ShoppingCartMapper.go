package mapper

import (
	"log"
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/commonParams"
	"strings"
)

type ShoppingCartMapper struct{}

func (s *ShoppingCartMapper) List(cart entity.ShoppingCart) (carts []entity.ShoppingCart, err error) {
	carts = make([]entity.ShoppingCart, 0)
	selectSQL := "select * from shopping_cart where true"
	args := []interface{}{}
	if cart.UserID != -1 {
		selectSQL = selectSQL + " and user_id = ?"
		args = append(args, cart.UserID)
	}
	if cart.SetmealID != -1 {
		selectSQL = selectSQL + " and setmeal_id = ?"
		args = append(args, cart.SetmealID)
	}
	if cart.DishID != -1 {
		selectSQL = selectSQL + " and dish_id = ?"
		args = append(args, cart.DishID)
	}
	if cart.DishFlavor != "" {
		selectSQL = selectSQL + " and dish_flavor = ?"
		args = append(args, cart.DishFlavor)
	}
	log.Println(selectSQL, args)
	rows, err := commonParams.Db.Query(selectSQL, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var row = entity.ShoppingCart{}
		err = rows.Scan(&row.ID, &row.Name, &row.Image, &row.UserID, &row.DishID, &row.SetmealID, &row.DishFlavor, &row.Number, &row.Amount, &row.CreateTime)
		if err != nil {
			return nil, err
		}
		carts = append(carts, row)
	}
	return carts, err

}

func (s *ShoppingCartMapper) UpdateNumberById(cart entity.ShoppingCart) error {
	updateSQL := "update shopping_cart set number = ? where id = ?"
	args := []interface{}{cart.Number, cart.ID}
	log.Println(updateSQL, args)
	_, err := commonParams.Db.Exec(updateSQL, args...)
	return err
}

func (s *ShoppingCartMapper) Insert(cart entity.ShoppingCart) error {
	insertSQL := "insert into shopping_cart (name, image, user_id, dish_id, setmeal_id, dish_flavor, amount, create_time) values (?,?,?,?,?,?,?,?)"
	args := []interface{}{cart.Name, cart.Image, cart.UserID, cart.DishID, cart.SetmealID, cart.DishFlavor, cart.Amount, cart.CreateTime}
	log.Println(insertSQL, args)
	_, err := commonParams.Db.Exec(insertSQL, args...)
	return err
}

func (s *ShoppingCartMapper) DeleteByUserId(userId int64) error {
	deleteSQL := "delete from shopping_cart where user_id = ?"
	args := []interface{}{userId}
	log.Println(deleteSQL, args)
	_, err := commonParams.Tx.Exec(deleteSQL, args...)
	return err
}

func (s *ShoppingCartMapper) InsertBatch(list []entity.ShoppingCart) error {
	insertSQL := "insert into shopping_cart (name, image, user_id, dish_id, setmeal_id, dish_flavor, amount,number, create_time) values "
	args := []interface{}{}
	for _, cart := range list {
		insertSQL = insertSQL + "(?, ?, ?, ?, ?, ?, ?, ?, ?),"
		args = append(args, cart.Name, cart.Image, cart.UserID, cart.DishID, cart.SetmealID, cart.DishFlavor, cart.Amount, cart.Number, cart.CreateTime)

	}
	insertSQL = strings.TrimSuffix(insertSQL, ",")
	log.Println(insertSQL, args)
	_, err := commonParams.Tx.Exec(insertSQL, args...)
	return err

}

func (s *ShoppingCartMapper) DeleteById(id int64) error {
	deleteSQL := "delete from shopping_cart where id = ?"
	args := []interface{}{id}
	log.Println(deleteSQL, args)
	_, err := commonParams.Db.Exec(deleteSQL, args...)
	return err
}
