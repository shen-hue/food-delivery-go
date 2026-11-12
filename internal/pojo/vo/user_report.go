package vo

// UserReportVO 用户统计返回结果（各列表为逗号分隔的字符串）。
type UserReportVO struct {
	DateList      string `json:"dateList"`      // 日期列表
	TotalUserList string `json:"totalUserList"` // 用户总量列表
	NewUserList   string `json:"newUserList"`   // 新增用户列表
}
