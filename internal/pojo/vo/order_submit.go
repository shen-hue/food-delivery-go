package vo

import "time"

// OrderSubmitVO 用户下单返回结果。
type OrderSubmitVO struct {
	ID          int64     `json:"id"`
	OrderNumber string    `json:"orderNumber"`
	OrderAmount float64   `json:"orderAmount"`
	OrderTime   time.Time `json:"orderTime"`
}
