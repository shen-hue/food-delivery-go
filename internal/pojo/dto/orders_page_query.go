package dto

import "sky-take-out/internal/common/jsonutil"

// OrdersPageQueryDTO 订单分页查询的数据传输对象（管理端条件查询 & 用户端历史订单查询）。
type OrdersPageQueryDTO struct {
	Page      int
	PageSize  int
	Number    string
	Phone     string
	Status    *int
	BeginTime jsonutil.DateTime // yyyy-MM-dd HH:mm:ss
	EndTime   jsonutil.DateTime // yyyy-MM-dd HH:mm:ss
	UserID    int64
}
