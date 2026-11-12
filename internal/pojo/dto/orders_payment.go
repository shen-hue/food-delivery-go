package dto

// OrdersPaymentDTO 订单支付的数据传输对象。
type OrdersPaymentDTO struct {
	OrderNumber string `json:"orderNumber"`
	PayMethod   *int   `json:"payMethod"`
}
