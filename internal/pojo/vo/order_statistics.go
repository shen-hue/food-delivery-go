package vo

// OrderStatisticsVO 各状态订单数量统计（管理端订单统计）。
type OrderStatisticsVO struct {
	ToBeConfirmed      int `json:"toBeConfirmed"`      // 待接单数量
	Confirmed          int `json:"confirmed"`          // 已接单数量
	DeliveryInProgress int `json:"deliveryInProgress"` // 派送中数量
}
