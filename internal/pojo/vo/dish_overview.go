package vo

// DishOverViewVO 菜品总览数据（工作台展示）。
type DishOverViewVO struct {
	Sold         int `json:"sold"`         // 已售数量
	Discontinued int `json:"discontinued"` // 停售数量
}
