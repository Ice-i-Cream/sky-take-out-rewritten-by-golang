package mapper

import (
	"log"
	"sky-take-out/common/result"
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/entity"
	"sky-take-out/pojo/vo"
	"sky-take-out/resources/commonParams"
	"sky-take-out/resources/functionParams"
	"strings"
	"time"
)

type DishMapper struct{}

func (d *DishMapper) CountByCategoryId(value int) (count int, err error) {
	selectSQL := "select COUNT(id) from Dish where category_id=?"
	log.Println(selectSQL, value)
	err = commonParams.Db.QueryRow(selectSQL, value).Scan(&count)
	return count, err
}

func (d *DishMapper) PageQuery(dto dto.DishPageQueryDTO) (res result.PageResult, err error) {
	res.Records = []interface{}{}
	selectSQL := "select d.*,c.name as categoryName from dish d left join category c on d.category_id = c.id where true"
	var args []interface{}
	if dto.Name != "" {
		selectSQL += " and d.name like concat('%',?,'%')"
		args = append(args, dto.Name)
	}
	if dto.CategoryID != -1 {
		selectSQL += " and d.category_id = ?"
		args = append(args, dto.CategoryID)
	}
	if dto.Status != -1 {
		selectSQL += " and d.status = ?"
		args = append(args, dto.Status)
	}
	selectSQL += " order by d.create_time desc limit ? offset ?" + ""
	args = append(args, dto.PageSize)
	args = append(args, (dto.Page-1)*dto.PageSize)
	log.Println(selectSQL, args)
	rows, err := commonParams.Db.Query(selectSQL, args...)
	if err != nil {
		return result.PageResult{}, err
	}
	res.Total = 0
	for rows.Next() {
		var dishVO vo.DishVO
		var ignore interface{}
		var err = rows.Scan(&dishVO.ID, &dishVO.Name, &dishVO.CategoryID, &dishVO.Price, &dishVO.Image, &dishVO.Description, &dishVO.Status, &ignore, &dishVO.UpdateTime, &ignore, &ignore, &dishVO.CategoryName)
		dishVO.Time = dishVO.UpdateTime.Format("2006-01-02 15:04:05")
		if err != nil {
			return res, err
		}
		res.Records = append(res.Records, dishVO)
		res.Total++
	}
	return res, err

}

func (d *DishMapper) Insert(dish entity.Dish) (dishId int64, err error) {
	insertSQL := "insert into dish (name, category_id, price, image, description, status, create_time, update_time, create_user, update_user) VALUES (?,?,?,?,?,?,?,?,?,?)"
	args := []interface{}{dish.Name, dish.CategoryID, dish.Price, dish.Image, dish.Description, dish.Status, dish.CreateTime, dish.UpdateTime, dish.CreateUser, dish.UpdateUser}
	log.Println(insertSQL, args)
	res, err := commonParams.Tx.Exec(insertSQL, args...)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (d *DishMapper) GetById(id string) (dish entity.Dish, err error) {
	selectSQL := "select * from dish where id=?"
	log.Println(selectSQL, id)
	rows, err := commonParams.Db.Query(selectSQL, id)
	if err != nil {
		return dish, err
	}
	rows.Next()
	err = rows.Scan(&dish.ID, &dish.Name, &dish.CategoryID, &dish.Price, &dish.Image, &dish.Description, &dish.Status, &dish.CreateTime, &dish.UpdateTime, &dish.CreateUser, &dish.UpdateUser)
	return dish, err
}

func (d *DishMapper) DeleteByIds(list []string) error {
	placeholders := make([]string, len(list))
	for i := range placeholders {
		placeholders[i] = "?"
	}
	placeholderStr := strings.Join(placeholders, ", ")

	var args []interface{}
	for _, id := range list {
		args = append(args, functionParams.ToInt(id))
	}

	deleteSQL := "delete from dish where id in ( " + placeholderStr + " ) "
	log.Println(deleteSQL, args)
	_, err := commonParams.Tx.Exec(deleteSQL, args...)
	return err
}

func (d *DishMapper) Update(dish entity.Dish) error {
	updateSQL := "update dish set status=?"
	var args []interface{}
	args = append(args, dish.Status)

	if dish.Name != "" {
		updateSQL += ", name=?"
		args = append(args, dish.Name)
	}
	if dish.CategoryID != -1 {
		updateSQL += ", category_id=?"
		args = append(args, dish.CategoryID)
	}
	if dish.Price != -1 {
		updateSQL += ", price=?"
		args = append(args, dish.Price)
	}
	if dish.Image != "" {
		updateSQL += ", image=?"
		args = append(args, dish.Image)
	}
	if dish.Description != "" {
		updateSQL += ", description=?"
		args = append(args, dish.Description)
	}
	if dish.UpdateTime != *new(time.Time) {
		updateSQL += ", update_time=?"
		args = append(args, dish.UpdateTime)
	}
	if dish.UpdateUser != -1 {
		updateSQL += ", update_user=?"
		args = append(args, dish.UpdateUser)
	}
	updateSQL = updateSQL + " where id=?"
	args = append(args, dish.ID)
	log.Println(updateSQL, args)
	_, err := commonParams.Db.Exec(updateSQL, args...)
	return err
}

func (d *DishMapper) List(dish entity.Dish) (list []entity.Dish, err error) {
	selectSQL := "select * from dish where true"
	args := []interface{}{}
	if dish.Name != "" {
		selectSQL += " and name like concat( '%' ? '%' )"
		args = append(args, dish.Name)
	}
	if dish.CategoryID != -1 {
		selectSQL += " and category_id=?"
		args = append(args, dish.CategoryID)
	}
	if dish.Status != -1 {
		selectSQL += " and status=?"
		args = append(args, dish.Status)
	}
	selectSQL = selectSQL + " order by create_time desc"
	log.Println(selectSQL, args)
	rows, err := commonParams.Db.Query(selectSQL, args...)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var dish entity.Dish
		err := rows.Scan(&dish.ID, &dish.Name, &dish.CategoryID, &dish.Price, &dish.Image, &dish.Description, &dish.Status, &dish.CreateTime, &dish.UpdateTime, &dish.CreateUser, &dish.UpdateUser)
		if err != nil {
			return nil, err
		}
		dish.CTime = dish.CreateTime.Format("2006-01-02 15:04:05")
		dish.UTime = dish.UpdateTime.Format("2006-01-02 15:04:05")

		list = append(list, dish)
	}
	return list, err
}

func (d *DishMapper) CountByMap(m map[interface{}]interface{}) (int64, error) {
	selectSQL := "select count(id) as count from dish where true"
	args := []interface{}{}
	status := m["status"].(int)
	categoryId := m["categoryId"].(int)
	if status != -1 {
		selectSQL += " and status = ?"
		args = append(args, status)
	}
	if categoryId != -1 {
		selectSQL += " and categoryId = ?"
		args = append(args, categoryId)
	}
	log.Println(selectSQL, args)
	rows, err := commonParams.Db.Query(selectSQL, args...)
	if err != nil {
		return 0, err
	}
	var count int64
	for rows.Next() {
		err = rows.Scan(&count)
	}
	return count, err
}
