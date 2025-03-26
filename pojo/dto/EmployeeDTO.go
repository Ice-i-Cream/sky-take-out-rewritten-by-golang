package dto

type EmployeeDTO struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Sex      string `json:"sex"`
	IDNumber string `json:"idNumber"` // 在Go中，通常使用下划线来分隔多个单词以提高可读性
}
