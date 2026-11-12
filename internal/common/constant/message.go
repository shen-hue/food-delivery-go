// Package constant 对应 Java 版 sky-common 的 com.sky.constant 包。
package constant

// 信息提示常量（对应 MessageConstant）
const (
	PasswordError              = "密码错误"            // PASSWORD_ERROR
	AccountNotFound            = "账号不存在"           // ACCOUNT_NOT_FOUND
	AccountLocked              = "账号被锁定"           // ACCOUNT_LOCKED
	UnknownError               = "未知错误"            // UNKNOWN_ERROR
	UserNotLogin               = "用户未登录"           // USER_NOT_LOGIN
	CategoryBeRelatedBySetmeal = "当前分类关联了套餐,不能删除"  // CATEGORY_BE_RELATED_BY_SETMEAL
	CategoryBeRelatedByDish    = "当前分类关联了菜品,不能删除"  // CATEGORY_BE_RELATED_BY_DISH
	ShoppingCartIsNull         = "购物车数据为空，不能下单"    // SHOPPING_CART_IS_NULL
	AddressBookIsNull          = "用户地址为空，不能下单"     // ADDRESS_BOOK_IS_NULL
	LoginFailed                = "登录失败"            // LOGIN_FAILED
	UploadFailed               = "文件上传失败"          // UPLOAD_FAILED
	SetmealEnableFailed        = "套餐内包含未启售菜品，无法启售" // SETMEAL_ENABLE_FAILED
	PasswordEditFailed         = "密码修改失败"          // PASSWORD_EDIT_FAILED
	DishOnSale                 = "起售中的菜品不能删除"      // DISH_ON_SALE
	SetmealOnSale              = "起售中的套餐不能删除"      // SETMEAL_ON_SALE
	DishBeRelatedBySetmeal     = "当前菜品关联了套餐,不能删除"  // DISH_BE_RELATED_BY_SETMEAL
	OrderStatusError           = "订单状态错误"          // ORDER_STATUS_ERROR
	OrderNotFound              = "订单不存在"           // ORDER_NOT_FOUND
)
