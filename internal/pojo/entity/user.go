package entity

import "time"

// User 用户实体类（对应 Java 版 User，微信用户）。
// 注意：PostgreSQL 中 user 是保留字，表名使用 users。
type User struct {
	ID         int64     `json:"id" db:"id"`
	Openid     string    `json:"openid" db:"openid"`
	Name       *string   `json:"name" db:"name"`
	Phone      *string   `json:"phone" db:"phone"`
	Sex        *string   `json:"sex" db:"sex"`
	IDNumber   *string   `json:"idNumber" db:"id_number"`
	Avatar     *string   `json:"avatar" db:"avatar"`
	CreateTime time.Time `json:"createTime" db:"create_time"`
}
