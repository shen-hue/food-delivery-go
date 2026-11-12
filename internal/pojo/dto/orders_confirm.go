package dto

// OrdersConfirmDTO 订单确认的数据传输对象。
type OrdersConfirmDTO struct {
	ID     int64 `json:"id"`
	Status *int  `json:"status"`
}
