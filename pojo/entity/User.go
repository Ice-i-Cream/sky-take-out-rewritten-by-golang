package entity

import (
	"time"
)

type User struct {
	ID         int64     `json:"id"`         // 用户ID
	OpenID     string    `json:"openid"`     // 微信用户唯一标识
	Name       string    `json:"name"`       // 姓名
	Phone      string    `json:"phone"`      // 手机号
	Sex        string    `json:"sex"`        // 性别 0 女 1 男
	IdNumber   string    `json:"idNumber"`   // 身份证号
	Avatar     string    `json:"avatar"`     // 头像
	CreateTime time.Time `json:"createTime"` // 注册时间
}
