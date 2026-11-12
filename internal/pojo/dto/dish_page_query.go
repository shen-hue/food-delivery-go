package dto

// DishPageQueryDTO 菜品分页查询的数据传输对象。
type DishPageQueryDTO struct {
	Page       int
	PageSize   int
	Name       string
	CategoryID int64
	Status     *int
}
