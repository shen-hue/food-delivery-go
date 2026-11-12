package entity

import "time"

// Dish 菜品实体类（对应 Java 版 Dish）。status: 0停售 1起售。
type Dish struct {
	ID          int64     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	CategoryID  int64     `json:"categoryId" db:"category_id"`
	Price       float64   `json:"price" db:"price"`
	Image       *string   `json:"image" db:"image"`
	Description *string   `json:"description" db:"description"`
	Status      int       `json:"status" db:"status"`
	CreateTime  time.Time `json:"createTime" db:"create_time"`
	UpdateTime  time.Time `json:"updateTime" db:"update_time"`
	CreateUser  *int64    `json:"createUser" db:"create_user"`
	UpdateUser  *int64    `json:"updateUser" db:"update_user"`
}
