package entity

import "time"

// Category 分类实体类（对应 Java 版 Category）。
// type: 1菜品分类 2套餐分类；status: 0禁用 1启用。
type Category struct {
	ID         int64     `json:"id" db:"id"`
	Type       int       `json:"type" db:"type"`
	Name       string    `json:"name" db:"name"`
	Sort       int       `json:"sort" db:"sort"`
	Status     int       `json:"status" db:"status"`
	CreateTime time.Time `json:"createTime" db:"create_time"`
	UpdateTime time.Time `json:"updateTime" db:"update_time"`
	CreateUser *int64    `json:"createUser" db:"create_user"`
	UpdateUser *int64    `json:"updateUser" db:"update_user"`
}
