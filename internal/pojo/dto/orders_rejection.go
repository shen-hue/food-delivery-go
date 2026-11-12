package dto

// OrdersRejectionDTO 订单拒绝的数据传输对象。
type OrdersRejectionDTO struct {
	ID              int64  `json:"id"`
	RejectionReason string `json:"rejectionReason"`
}
