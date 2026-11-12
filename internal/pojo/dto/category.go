package dto

// CategoryDTO 新增/编辑分类时的数据传输对象。
type CategoryDTO struct {
	ID   int64  `json:"id"`
	Type int    `json:"type"` // 1菜品分类 2套餐分类
	Name string `json:"name"`
	Sort int    `json:"sort"`
}
