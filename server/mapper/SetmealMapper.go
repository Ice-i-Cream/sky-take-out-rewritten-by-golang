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

func (s *SetmealMapper) PageQuery(dto dto.SetmealPageQueryDTO) (res result.PageResult, err error) {
	selectSQL := "  select s.*,c.name categoryName from setmeal s left join category c on s.category_id = c.id where true"
	var args []interface{}
	if dto.Name != "" {
		selectSQL += " and s.name like concat ('%' ,? ,'%')"
		args = append(args, dto.Name)
	}
	if dto.CategoryId != -1 {
		selectSQL += " and s.category_id = ?"
		args = append(args, dto.CategoryId)
	}
	if dto.Status != -1 {
		selectSQL += " and s.status = ?"
		args = append(args, dto.Status)
	}
	selectSQL += " order by s.create_time desc limit ? offset ?" + ""
	args = append(args, dto.PageSize)
	args = append(args, (dto.Page-1)*dto.PageSize)
	log.Println(selectSQL, args)
	rows, err := commonParams.Db.Query(selectSQL, args...)
	if err != nil {
		return res, err
	}
	setmealVO := vo.SetmealVO{}
	var ignore interface{}
	res.Total = 0
	res.Records = []interface{}{}
	for rows.Next() {
		err = rows.Scan(&setmealVO.ID, &setmealVO.CategoryID, &setmealVO.Name, &setmealVO.Price, &setmealVO.Status, &setmealVO.Description, &setmealVO.Image, &ignore, &setmealVO.UpdateTime, &ignore, &ignore, &setmealVO.CategoryName)
		setmealVO.UTime = setmealVO.UpdateTime.Format("2006-01-02 15:04:05")
		setmealVO.SetmealDishes = []entity.SetmealDish{}
		res.Records = append(res.Records, setmealVO)
		res.Total++
	}
	return res, nil

}

func (s *SetmealMapper) Insert(setmeal entity.Setmeal) (int, error) {
	insertSQL := "insert into setmeal (category_id, name, price, description, image, create_time, update_time, create_user, update_user) VALUES (?,?,?,?,?,?,?,?,?)"
	args := []interface{}{setmeal.CategoryId, setmeal.Name, setmeal.Price, setmeal.Description, setmeal.Image, setmeal.CreateTime, setmeal.UpdateTime, setmeal.CreateUser, setmeal.UpdateUser}
	exec, err := functionParams.ExecSQL(insertSQL, args)
	if err != nil {
		return 0, err
	}
	id, _ := exec.LastInsertId()
	return int(id), err
}

func (s *SetmealMapper) GetById(id int) (setmeal entity.Setmeal, err error) {
	selectSQL := "select * from setmeal where id = ?"
	exec, err := commonParams.Db.Query(selectSQL, id)
	if err != nil {
		return setmeal, err
	}
	exec.Next()
	err = exec.Scan(&setmeal.Id, &setmeal.CategoryId, &setmeal.Name, &setmeal.Price, &setmeal.Status, &setmeal.Description, &setmeal.Image, &setmeal.CreateTime, &setmeal.UpdateTime, &setmeal.CreateUser, &setmeal.UpdateUser)
	return setmeal, err
}

func (s *SetmealMapper) DeleteById(id int) error {
	deleteSQL := "delete from setmeal where id = ?"
	_, err := functionParams.ExecSQL(deleteSQL, []interface{}{id})
	return err

}

func (s *SetmealMapper) Update(setmeal entity.Setmeal) error {
	updateSQL := "update setmeal set status = ?"
	args := []interface{}{setmeal.Status}

	if setmeal.Name != "" {
		updateSQL += ", name = ?"
		args = append(args, setmeal.Name)
	}
	if setmeal.CategoryId != -1 {
		updateSQL += ", category_id = ?"
		args = append(args, setmeal.CategoryId)
	}
	if setmeal.Price != -1 {
		updateSQL += ", price = ?"
		args = append(args, setmeal.Price)
	}
	if setmeal.Description != "" {
		updateSQL += ", description = ?"
		args = append(args, setmeal.Description)
	}
	if setmeal.Image != "" {
		updateSQL += ", image = ?"
		args = append(args, setmeal.Image)
	}
	if setmeal.UpdateTime != *new(time.Time) {
		updateSQL += ", update_time = ?"
		args = append(args, setmeal.UpdateTime)
	}
	if setmeal.UpdateUser != -1 {
		updateSQL += ", update_user = ?"
		args = append(args, setmeal.UpdateUser)
	}
	updateSQL += " where id = ?"
	args = append(args, setmeal.Id)

	log.Println(updateSQL, args)
	_, err := functionParams.ExecSQL(updateSQL, args)
	return err
}

func (s *SetmealMapper) List(setmeal entity.Setmeal) (setmeals []entity.Setmeal, err error) {
	selectSQL := "select * from setmeal where true"
	args := []interface{}{}

	setmeals = []entity.Setmeal{}
	if setmeal.Name != "" {
		selectSQL += " and name = ?"
		args = append(args, setmeal.Name)
	}
	if setmeal.CategoryId != -1 {
		selectSQL += " and category_id = ?"
		args = append(args, setmeal.CategoryId)
	}
	if setmeal.Status != -1 {
		selectSQL += " and status = ?"
		args = append(args, setmeal.Status)
	}
	log.Println(selectSQL, args)
	rows, err := commonParams.Db.Query(selectSQL, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var setmeal entity.Setmeal
		_ = rows.Scan(&setmeal.Id, &setmeal.CategoryId, &setmeal.Name, &setmeal.Price, &setmeal.Status, &setmeal.Description, &setmeal.Image, &setmeal.CreateTime, &setmeal.UpdateTime, &setmeal.CreateUser, &setmeal.UpdateUser)
		setmeals = append(setmeals, setmeal)
	}
	return setmeals, err

}

func (s *SetmealMapper) GetDishItemBySetmealId(id int64) (vos []vo.DishItemVO, err error) {
	selectSQL := "select sd.name, sd.copies, d.image, d.description from setmeal_dish sd left join dish d on sd.dish_id = d.id where sd.setmeal_id = ?"
	args := []interface{}{id}
	log.Println(selectSQL, args)
	rows, err := commonParams.Db.Query(selectSQL, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var dishvo = vo.DishItemVO{}
		err := rows.Scan(&dishvo.Name, &dishvo.Copies, &dishvo.Image, &dishvo.Description)
		if err != nil {
			return nil, err
		}
		vos = append(vos, dishvo)
	}
	return vos, err
}

func (s *SetmealMapper) CountByMap(m map[interface{}]interface{}) (int64, error) {
	selectSQL := "select count(id) as count from setmeal where true"
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
