// Package dto 对应 Java 版 sky-pojo 的 com.sky.dto 包。
package dto

// EmployeeLoginDTO 员工登录时的数据传输对象。
type EmployeeLoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
