package dto

import "sky-take-out/internal/pojo/entity"

// DishDTO 新增/编辑菜品时的数据传输对象（含口味列表）。
type DishDTO struct {
	ID          int64               `json:"id"`
	Name        string              `json:"name"`
	CategoryID  int64               `json:"categoryId"`
	Price       float64             `json:"price"`
	Image       string              `json:"image"`
	Description string              `json:"description"`
	Status      int                 `json:"status"`
	Flavors     []entity.DishFlavor `json:"flavors"`
}
