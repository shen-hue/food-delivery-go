package vo

// OrderOverViewVO 订单总览数据（工作台展示）。
type OrderOverViewVO struct {
	WaitingOrders   int `json:"waitingOrders"`   // 待接单数量
	DeliveredOrders int `json:"deliveredOrders"` // 派送中数量
	CompletedOrders int `json:"completedOrders"` // 已完成数量
	CancelledOrders int `json:"cancelledOrders"` // 已取消数量
	AllOrders       int `json:"allOrders"`       // 全部订单数量
}
