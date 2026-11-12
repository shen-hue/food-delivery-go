package vo

// DishItemVO 套餐内菜品展示对象（用户端浏览套餐明细时使用）。
type DishItemVO struct {
	Name        string `json:"name"`
	Copies      int    `json:"copies"`
	Image       string `json:"image"`
	Description string `json:"description"`
}
