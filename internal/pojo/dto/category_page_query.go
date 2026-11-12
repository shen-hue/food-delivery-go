package dto

// CategoryPageQueryDTO 分类分页查询的数据传输对象。
type CategoryPageQueryDTO struct {
	Page     int
	PageSize int
	Name     string
	Type     *int
}
