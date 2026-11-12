// Package service 对应 Java 版 sky-server 的 com.sky.service 包。
package service

import (
	"context"

	"sky-take-out/internal/pojo/dto"
	"sky-take-out/internal/pojo/entity"
)

// EmployeeService 员工业务接口（对应 Java 版 EmployeeService 接口）。
type EmployeeService interface {
	// Login 员工登录
	Login(ctx context.Context, employeeLoginDTO *dto.EmployeeLoginDTO) (*entity.Employee, error)
}
