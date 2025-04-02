package entity

import "time"

type Category struct {
	ID         int64     `json:"id"`
	Type       int       `json:"type"`
	Name       string    `json:"name"`
	Sort       int       `json:"sort"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"-"`
	CreateUser int64     `json:"createUser"`
	UpdateUser int64     `json:"updateUser"`
	Time       string    `json:"updateTime"`
}
