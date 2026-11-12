package vo

// BusinessDataVO 工作台今日数据展示对象。
type BusinessDataVO struct {
	Turnover            float64 `json:"turnover"`            // 营业额
	ValidOrderCount     int     `json:"validOrderCount"`     // 有效订单数
	OrderCompletionRate float64 `json:"orderCompletionRate"` // 订单完成率
	UnitPrice           float64 `json:"unitPrice"`           // 平均客单价
	NewUsers            int     `json:"newUsers"`            // 新增用户数
}
