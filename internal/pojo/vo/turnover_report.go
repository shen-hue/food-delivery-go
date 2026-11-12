package vo

// TurnoverReportVO 营业额统计返回结果（dateList/turnoverList为逗号分隔的字符串）。
type TurnoverReportVO struct {
	DateList     string `json:"dateList"`     // 日期列表，如 "2026-11-01,2026-11-02"
	TurnoverList string `json:"turnoverList"` // 营业额列表，如 "100.0,200.0"
}
