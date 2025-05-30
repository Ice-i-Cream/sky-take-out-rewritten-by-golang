package mapper

import (
	"log"
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/commonParams"
	"time"
)

type UserMapper struct{}

func (u *UserMapper) GetByOpenid(openid string) (entity.User, error) {
	selectSQL := "select * from user where openid=?"
	log.Println(selectSQL, openid)
	rows, err := commonParams.Db.Query(selectSQL, openid)
	if err != nil {
		return entity.User{}, err
	}
	rows.Next()
	row := entity.User{}
	err = rows.Scan(&row.ID, &row.OpenID, &row.Name, &row.Phone, &row.Sex, &row.IdNumber, &row.Avatar, &row.CreateTime)
	if err != nil {
		return entity.User{}, err
	}
	return row, nil
}

func (u *UserMapper) Insert(user entity.User) error {
	insertSQL := "insert into user (openid, name, phone, sex, id_number, avatar, create_time) values (?,?,?,?,?,?,?)"
	args := []interface{}{user.OpenID, user.Name, user.Phone, user.Sex, user.IdNumber, user.Avatar, user.CreateTime}
	log.Println(insertSQL, args)
	_, err := commonParams.Tx.Exec(insertSQL, args...)
	return err
}

func (u *UserMapper) CountByMap(m map[interface{}]interface{}) (int64, error) {
	begin := m["begin"].(time.Time)
	end := m["end"].(time.Time)
	selectSQL := "select count(id) as count from user where true"
	args := []interface{}{}
	if begin != *new(time.Time) {
		selectSQL = selectSQL + " and create_time >= ?"
		args = append(args, begin)
	}
	if end != *new(time.Time) {
		selectSQL = selectSQL + " and create_time < ?"
		args = append(args, end)
	}
	//log.Println(selectSQL, args)
	rows, err := commonParams.Db.Query(selectSQL, args...)
	if err != nil {
		return 0, err
	}
	var count int64
	for rows.Next() {
		err = rows.Scan(&count)
	}
	return count, nil
}
