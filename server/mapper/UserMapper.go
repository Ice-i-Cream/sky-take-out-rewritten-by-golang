package mapper

import (
	"log"
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/commonParams"
)

type UserMapper struct{}

func (m UserMapper) GetByOpenid(openid string) (entity.User, error) {
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

func (m UserMapper) Insert(user entity.User) error {
	insertSQL := "insert into user (openid, name, phone, sex, id_number, avatar, create_time) values (?,?,?,?,?,?,?)"
	args := []interface{}{user.OpenID, user.Name, user.Phone, user.Sex, user.IdNumber, user.Avatar, user.CreateTime}
	log.Println(insertSQL, args)
	_, err := commonParams.Tx.Exec(insertSQL, args...)
	return err
}
