// Package impl 对应 Java 版 sky-server 的 com.sky.service.impl 包。
package impl

import (
	"context"

	"sky-take-out/internal/common/constant"
	"sky-take-out/internal/common/exception"
	"sky-take-out/internal/pojo/dto"
	"sky-take-out/internal/pojo/entity"
	"sky-take-out/internal/server/mapper"
	"sky-take-out/internal/server/service"
)

// EmployeeServiceImpl 员工业务实现类（对应 Java 版 EmployeeServiceImpl，@Service）。
type EmployeeServiceImpl struct {
	employeeMapper *mapper.EmployeeMapper
}

// 编译期检查：确保实现了 EmployeeService 接口（相当于 Java 的 implements 关键字）。
var _ service.EmployeeService = (*EmployeeServiceImpl)(nil)

// NewEmployeeServiceImpl 构造函数（对应 @Autowired 依赖注入）。
func NewEmployeeServiceImpl(employeeMapper *mapper.EmployeeMapper) *EmployeeServiceImpl {
	return &EmployeeServiceImpl{employeeMapper: employeeMapper}
}

// Login 员工登录。
//
// 对应 Java 版 EmployeeServiceImpl.login：
//  1. 根据用户名查询数据库中的数据
//  2. 处理各种异常情况（用户名不存在、密码不对、账号被锁定）
//  3. 返回实体对象
func (s *EmployeeServiceImpl) Login(ctx context.Context, employeeLoginDTO *dto.EmployeeLoginDTO) (*entity.Employee, error) {
	username := employeeLoginDTO.Username
	password := employeeLoginDTO.Password

	// 1、根据用户名查询数据库中的数据
	employee, err := s.employeeMapper.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	// 2、处理各种异常情况（用户名不存在、密码不对、账号被锁定）
	if employee == nil {
		// 账号不存在
		return nil, exception.NewAccountNotFoundException(constant.AccountNotFound)
	}

	// 密码比对
	// TODO 后期需要进行md5加密，然后再进行比对
	// 练习提示：Go 中可用 crypto/md5 包，
	// hex.EncodeToString(md5.Sum([]byte(password))[:]) 得到摘要后再比对
	if password != employee.Password {
		// 密码错误
		return nil, exception.NewPasswordErrorException(constant.PasswordError)
	}

	if employee.Status == constant.Disable {
		// 账号被锁定
		return nil, exception.NewAccountLockedException(constant.AccountLocked)
	}

	// 3、返回实体对象
	return employee, nil
}
