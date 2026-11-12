package entity

import "time"

// Setmeal 套餐实体类（对应 Java 版 Setmeal）。status: 0停用 1启用。
type Setmeal struct {
	ID          int64     `json:"id" db:"id"`
	CategoryID  int64     `json:"categoryId" db:"category_id"`
	Name        string    `json:"name" db:"name"`
	Price       float64   `json:"price" db:"price"`
	Status      int       `json:"status" db:"status"`
	Description *string   `json:"description" db:"description"`
	Image       *string   `json:"image" db:"image"`
	CreateTime  time.Time `json:"createTime" db:"create_time"`
	UpdateTime  time.Time `json:"updateTime" db:"update_time"`
	CreateUser  *int64    `json:"createUser" db:"create_user"`
	UpdateUser  *int64    `json:"updateUser" db:"update_user"`
}
