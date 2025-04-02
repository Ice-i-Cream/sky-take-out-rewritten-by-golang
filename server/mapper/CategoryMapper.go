package mapper

import (
	"log"
	"sky-take-out/common/result"
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/commonParams"
	"strings"
	"time"
)

type CategoryMapper struct{}

func (c *CategoryMapper) PageQuery(dto dto.CategoryPageQueryDTO) (res result.PageResult, err error) {
	selectSQL := "select * from category where true"
	var args []interface{}
	res.Records = []interface{}{}
	if dto.Name != "" {
		selectSQL = selectSQL + " and name like concat('%', ? ,'%')"
		args = append(args, dto.Name)
	}
	if dto.Type != -1 {
		selectSQL = selectSQL + " and type = ?"
		args = append(args, dto.Type)
	}
	args = append(args, dto.PageSize)
	args = append(args, (dto.Page-1)*dto.PageSize)
	selectSQL = selectSQL + " order by create_time desc limit ? offset ?"
	log.Println(selectSQL, args)
	rows, err := commonParams.Db.Query(selectSQL, args...)
	if err != nil {
		return res, err
	}
	res.Total = 0
	for rows.Next() {
		var category entity.Category
		err = rows.Scan(&category.ID, &category.Type, &category.Name, &category.Sort, &category.Status, &category.CreateTime, &category.UpdateTime, &category.CreateUser, &category.UpdateUser)
		category.Time = category.UpdateTime.Format("2006-01-02 15:04:05")
		if err != nil {
			return res, err
		}
		res.Records = append(res.Records, category)
		res.Total++
	}
	return res, err
}

func (c *CategoryMapper) Save(category entity.Category) error {
	insertSQL := "insert into category (type, name, sort, status, create_time, update_time, create_user, update_user) VALUES (?,?,?,?,?,?,?,?)"
	args := []interface{}{category.Type, category.Name, category.Sort, category.Status, category.CreateTime, category.UpdateTime, category.CreateUser, category.UpdateUser}
	log.Println(insertSQL, args)
	_, err := commonParams.Db.Exec(insertSQL, args...)
	return err
}

func (c *CategoryMapper) DeleteById(value int) error {
	deleteSQL := "delete from category where id=?"
	log.Println(deleteSQL, value)
	_, err := commonParams.Db.Exec(deleteSQL, value)
	return err
}

func (c *CategoryMapper) Update(category entity.Category) error {
	updateSQL := "update category set" + ""
	var args []interface{}
	var updates []string
	if category.Name != "" {
		updates = append(updates, " name=?")
		args = append(args, category.Name)
	}
	if category.Sort != -1 {
		updates = append(updates, " sort=?")
		args = append(args, category.Sort)
	}
	if category.Status != -1 {
		updates = append(updates, " status=?")
		args = append(args, category.Status)
	}
	if category.UpdateTime != *new(time.Time) {
		updates = append(updates, " update_time=?")
		args = append(args, category.UpdateTime)
	}
	if category.UpdateUser != -1 {
		updates = append(updates, " update_user=?")
		args = append(args, category.UpdateUser)
	}
	updateSQL += strings.Join(updates, ",")
	updateSQL += " where id=?" + ""
	args = append(args, category.ID)
	log.Println(updateSQL, args)
	_, err := commonParams.Db.Exec(updateSQL, args...)
	return err
}

func (c *CategoryMapper) List(kind int) (list []entity.Category, err error) {
	selectSQL := "select * from category where type=?"
	log.Println(selectSQL, kind)
	rows, err := commonParams.Db.Query(selectSQL, kind)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var category entity.Category
		err = rows.Scan(&category.ID, &category.Type, &category.Name, &category.Sort, &category.Status, &category.CreateTime, &category.UpdateTime, &category.CreateUser, &category.UpdateUser)
		category.Time = category.UpdateTime.Format("2006-01-02 15:04:05")
		if err != nil {
			return list, err
		}
		list = append(list, category)
	}
	return list, err
}
