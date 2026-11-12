package vo

import (
	"time"

	"sky-take-out/internal/pojo/entity"
)

// DishVO 菜品展示对象（含口味与分类名称）。
type DishVO struct {
	ID           int64               `json:"id"`
	Name         string              `json:"name"`
	CategoryID   int64               `json:"categoryId"`
	Price        float64             `json:"price"`
	Image        string              `json:"image"`
	Description  string              `json:"description"`
	Status       int                 `json:"status"`
	UpdateTime   time.Time           `json:"updateTime"`
	CategoryName string              `json:"categoryName"`
	Flavors      []entity.DishFlavor `json:"flavors"`
}
