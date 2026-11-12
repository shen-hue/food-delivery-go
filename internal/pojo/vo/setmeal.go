package vo

import (
	"time"

	"sky-take-out/internal/pojo/entity"
)

// SetmealVO 套餐展示对象（含关联菜品与分类名称）。
type SetmealVO struct {
	ID            int64                `json:"id"`
	CategoryID    int64                `json:"categoryId"`
	Name          string               `json:"name"`
	Price         float64              `json:"price"`
	Status        int                  `json:"status"`
	Description   string               `json:"description"`
	Image         string               `json:"image"`
	UpdateTime    time.Time            `json:"updateTime"`
	CategoryName  string               `json:"categoryName"`
	SetmealDishes []entity.SetmealDish `json:"setmealDishes"`
}
