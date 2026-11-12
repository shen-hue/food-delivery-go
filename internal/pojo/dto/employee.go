package dto

// EmployeeDTO 新增/编辑员工时的数据传输对象。
type EmployeeDTO struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Sex      string `json:"sex"`
	IDNumber string `json:"idNumber"`
}
