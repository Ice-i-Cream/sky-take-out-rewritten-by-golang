package mapper

import (
	"log"
	"sky-take-out/resources/commonParams"
	"sky-take-out/resources/functionParams"
	"strings"
)

type SetmealMapper struct{}

func (s *SetmealMapper) CountByCategoryId(value int) (count int, err error) {
	selectSQL := "select COUNT(id) from setmeal where category_id=?"
	log.Println(selectSQL, value)
	err = commonParams.Db.QueryRow(selectSQL, value).Scan(&count)
	return count, err
}

func (s *SetmealMapper) GetSetmealIdByDishIds(list []string) ([]string, error) {
	var args []interface{}
	for _, i := range list {
		args = append(args, functionParams.ToInt(i))
	}

	placeholders := make([]string, len(list))
	for i := range placeholders {
		placeholders[i] = "?"
	}
	placeholderStr := strings.Join(placeholders, ", ")

	selectSQL := "select setmeal_id from setmeal_dish where dish_id in ( " + placeholderStr + " )"
	log.Println(selectSQL, args)

	rows, err := commonParams.Db.Query(selectSQL, args...)
	if err != nil {
		return []string{}, err
	}
	list = []string{}
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			return []string{}, err
		}
		list = append(list, id)
	}
	return list, nil
}
