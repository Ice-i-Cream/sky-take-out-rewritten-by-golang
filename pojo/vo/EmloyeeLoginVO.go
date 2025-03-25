package vo

type EmployeeLoginVO struct {
	ID       int64  `json:"id"`       // 主键值
	UserName string `json:"userName"` // 用户名
	Name     string `json:"name"`     // 姓名
	Token    string `json:"token"`    // jwt令牌
}
