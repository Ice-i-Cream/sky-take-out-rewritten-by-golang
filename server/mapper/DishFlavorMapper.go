package mapper

import (
	"log"
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/commonParams"
	"sky-take-out/resources/functionParams"
	"strings"
)

type DishFlavorMapper struct{}

func (d *DishFlavorMapper) InsertBatch(list []entity.DishFlavor, dishId int64) error {
	insertSQL := "insert into dish_flavor (dish_id, name, value) values (?,?,?)"
	var args []interface{}
	for index, flavor := range list {
		if index != 0 {
			insertSQL = insertSQL + ",(?,?,?)"
		}
		args = append(args, dishId, flavor.Name, flavor.Value)
	}
	log.Println(insertSQL, args)
	_, err := functionParams.ExecSQL(insertSQL, args)
	return err
}

func (d *DishFlavorMapper) DeleteByDishIds(list []string) error {
	placeholders := make([]string, len(list))
	for i := range placeholders {
		placeholders[i] = "?"
	}
	placeholderStr := strings.Join(placeholders, ", ")

	var args []interface{}
	for _, dishId := range list {
		args = append(args, functionParams.ToInt(dishId))
	}

	deleteSQL := "delete from dish_flavor where dish_flavor.dish_id in ( " + placeholderStr + " ) "

	log.Println(deleteSQL, args)
	_, err := functionParams.ExecSQL(deleteSQL, args)
	return err
}

func (d *DishFlavorMapper) GetByDishId(id int) (list []entity.DishFlavor, err error) {
	selectSQL := "select * from dish_flavor where dish_flavor.dish_id = ?"
	rows, err := commonParams.Db.Query(selectSQL, id)
	if err != nil {
		return nil, err
	}
	var dishFlavor entity.DishFlavor
	for rows.Next() {
		if err = rows.Scan(&dishFlavor.ID, &dishFlavor.DishID, &dishFlavor.Name, &dishFlavor.Value); err != nil {
			return nil, err
		}
		list = append(list, dishFlavor)
	}
	return list, nil
}

func (d *DishFlavorMapper) DeleteByDishId(id int64) error {
	deleteSQL := "delete from dish_flavor where dish_id = ?"
	log.Println(deleteSQL, id)
	_, err := functionParams.ExecSQL(deleteSQL, []interface{}{id})
	return err

}
