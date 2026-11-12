package vo

// UserLoginVO 用户登录返回结果（微信用户）。
type UserLoginVO struct {
	ID     int64  `json:"id"`
	Openid string `json:"openid"`
	Token  string `json:"token"`
}
