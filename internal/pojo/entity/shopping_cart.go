package entity

import "time"

// ShoppingCart 购物车实体类（对应 Java 版 ShoppingCart）。
type ShoppingCart struct {
	ID         int64     `json:"id" db:"id"`
	Name       *string   `json:"name" db:"name"`
	UserID     int64     `json:"userId" db:"user_id"`
	DishID     *int64    `json:"dishId" db:"dish_id"`
	SetmealID  *int64    `json:"setmealId" db:"setmeal_id"`
	DishFlavor *string   `json:"dishFlavor" db:"dish_flavor"`
	Number     int       `json:"number" db:"number"`
	Amount     float64   `json:"amount" db:"amount"`
	Image      *string   `json:"image" db:"image"`
	CreateTime time.Time `json:"createTime" db:"create_time"`
}
