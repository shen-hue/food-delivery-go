package dto

// SetmealPageQueryDTO 套餐分页查询的数据传输对象。
type SetmealPageQueryDTO struct {
	Page       int
	PageSize   int
	Name       string
	CategoryID int64
	Status     *int
}
