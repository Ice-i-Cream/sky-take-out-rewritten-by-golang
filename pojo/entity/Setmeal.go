package entity

import "time"

type Setmeal struct {
	Id          int       `json:"id"`
	CategoryId  int64     `json:"categoryId"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Status      int       `json:"status"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
	CreateUser  int64     `json:"createUser"`
	UpdateUser  int64     `json:"updateUser"`
}
