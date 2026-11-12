package result

// PageResult 封装分页查询结果（对应 Java 版的 PageResult）。
type PageResult struct {
	Total   int64 `json:"total"`   // 总记录数
	Records any   `json:"records"` // 当前页数据集合
}
