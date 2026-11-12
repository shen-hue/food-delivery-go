package vo

// SalesTop10ReportVO 销量Top10统计返回结果（各列表为逗号分隔的字符串）。
type SalesTop10ReportVO struct {
	NameList   string `json:"nameList"`   // 商品名称列表
	NumberList string `json:"numberList"` // 销量列表
}
