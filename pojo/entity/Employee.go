package entity

import (
	"time"
)

// Employee 结构体定义
type Employee struct {
	ID         int64     `json:"id"`
	Username   string    `json:"username"`
	Name       string    `json:"name"`
	Password   string    `json:"-"`
	Phone      string    `json:"phone"`
	Sex        string    `json:"sex"`
	IDNumber   string    `json:"idNumber"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	CreateUser int64     `json:"createUser"`
	UpdateUser int64     `json:"updateUser"`
}
