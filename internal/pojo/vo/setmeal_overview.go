package vo

// SetmealOverViewVO 套餐总览数据（工作台展示）。
type SetmealOverViewVO struct {
	Sold         int `json:"sold"`         // 已售数量
	Discontinued int `json:"discontinued"` // 停售数量
}
