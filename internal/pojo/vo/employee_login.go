// Package vo 对应 Java 版 sky-pojo 的 com.sky.vo 包。
package vo

// EmployeeLoginVO 员工登录返回结果。
type EmployeeLoginVO struct {
	ID       int64  `json:"id"`
	UserName string `json:"userName"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}
