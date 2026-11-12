// Package exception 对应 Java 版的 com.sky.exception 包。
//
// Go 惯用法说明：Java 通过异常类继承体系（BaseException 及其子类）
// 表达业务异常；Go 中更惯用的做法是 error 值 + 具名的构造函数。
// 这里保留一个 BaseException 类型承载所有业务异常，
// 构造函数名与 Java 的异常类一一对应，语义不变：
//
//	exception.NewAccountNotFoundException(constant.AccountNotFound)
//	// 相当于 throw new AccountNotFoundException(MessageConstant.ACCOUNT_NOT_FOUND)
//
// 全局异常处理器（internal/server/handler）通过 errors.As 识别
// *BaseException 并统一转换为 Result.error 响应。
package exception

// BaseException 业务异常基类。
type BaseException struct {
	Msg string
}

func (e *BaseException) Error() string { return e.Msg }

// NewBaseException 对应 new BaseException(msg)。
func NewBaseException(msg string) *BaseException {
	return &BaseException{Msg: msg}
}

// ---- 以下构造函数对应 Java 版的各异常子类 ----

// NewAccountNotFoundException 账号不存在异常。
func NewAccountNotFoundException(msg string) *BaseException { return NewBaseException(msg) }

// NewAccountLockedException 账号被锁定异常。
func NewAccountLockedException(msg string) *BaseException { return NewBaseException(msg) }

// NewPasswordErrorException 密码错误异常。
func NewPasswordErrorException(msg string) *BaseException { return NewBaseException(msg) }

// NewLoginFailedException 登录失败异常。
func NewLoginFailedException(msg string) *BaseException { return NewBaseException(msg) }

// NewPasswordEditFailedException 密码修改失败异常。
func NewPasswordEditFailedException(msg string) *BaseException { return NewBaseException(msg) }

// NewUserNotLoginException 用户未登录异常。
func NewUserNotLoginException(msg string) *BaseException { return NewBaseException(msg) }

// NewDeletionNotAllowedException 删除不允许异常。
func NewDeletionNotAllowedException(msg string) *BaseException { return NewBaseException(msg) }

// NewSetmealEnableFailedException 套餐启用失败异常。
func NewSetmealEnableFailedException(msg string) *BaseException { return NewBaseException(msg) }

// NewOrderBusinessException 订单业务异常。
func NewOrderBusinessException(msg string) *BaseException { return NewBaseException(msg) }

// NewShoppingCartBusinessException 购物车业务异常。
func NewShoppingCartBusinessException(msg string) *BaseException { return NewBaseException(msg) }

// NewAddressBookBusinessException 地址簿业务异常。
func NewAddressBookBusinessException(msg string) *BaseException { return NewBaseException(msg) }
