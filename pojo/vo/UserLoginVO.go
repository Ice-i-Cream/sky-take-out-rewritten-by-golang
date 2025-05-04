package vo

type UserLoginVO struct {
	ID     int64  `json:"id"`
	OpenID string `json:"openid"`
	Token  string `json:"token"`
}
