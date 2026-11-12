package dto

// OrdersCancelDTO 订单取消的数据传输对象。
type OrdersCancelDTO struct {
	ID           int64  `json:"id"`
	CancelReason string `json:"cancelReason"`
}
