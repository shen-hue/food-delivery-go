package dto

// UserLoginDTO 用户登录的数据传输对象（code为微信登录授权码）。
type UserLoginDTO struct {
	Code string `json:"code"`
}
