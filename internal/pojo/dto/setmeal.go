package dto

import "sky-take-out/internal/pojo/entity"

// SetmealDTO 新增/编辑套餐时的数据传输对象（含关联菜品列表）。
type SetmealDTO struct {
	ID            int64                `json:"id"`
	CategoryID    int64                `json:"categoryId"`
	Name          string               `json:"name"`
	Price         float64              `json:"price"`
	Status        int                  `json:"status"`
	Description   string               `json:"description"`
	Image         string               `json:"image"`
	SetmealDishes []entity.SetmealDish `json:"setmealDishes"`
}
