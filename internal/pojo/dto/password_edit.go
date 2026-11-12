package dto

// PasswordEditDTO 修改密码的数据传输对象。
type PasswordEditDTO struct {
	EmpID       int64  `json:"empId"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
