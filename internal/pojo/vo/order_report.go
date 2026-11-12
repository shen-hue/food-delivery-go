package vo

// OrderReportVO 订单统计返回结果（各列表为逗号分隔的字符串）。
type OrderReportVO struct {
	DateList            string  `json:"dateList"`            // 日期列表
	OrderCountList      string  `json:"orderCountList"`      // 订单总数列表
	ValidOrderCountList string  `json:"validOrderCountList"` // 有效订单数列表
	TotalOrderCount     int     `json:"totalOrderCount"`     // 订单总数
	ValidOrderCount     int     `json:"validOrderCount"`     // 有效订单数
	OrderCompletionRate float64 `json:"orderCompletionRate"` // 订单完成率
}
