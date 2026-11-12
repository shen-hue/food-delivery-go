package vo

// OrderPaymentVO 订单支付返回结果（微信支付参数）。
type OrderPaymentVO struct {
	NonceStr   string `json:"nonceStr"`
	PaySign    string `json:"paySign"`
	TimeStamp  string `json:"timeStamp"`
	SignType   string `json:"signType"`
	PackageStr string `json:"package"`
}
