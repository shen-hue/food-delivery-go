// Package entity 对应 Java 版 sky-pojo 的 com.sky.entity 包。
//
// 说明：
//   - json 标签与 Java 版 Jackson 默认的 lowerCamelCase 输出一致
//   - db 标签（下划线列名）供 pgx 的 RowToStructByName 按 "select *"
//     自动扫描到结构体字段，相当于 MyBatis 的 map-underscore-to-camel-case
//   - Java 中可空的 Long/LocalDateTime 字段对应 Go 的 *int64/*time.Time，
//     为 NULL 时序列化为 null，与 Java 版行为一致
package entity

import "time"

// Employee 员工实体类（对应 Java 版 Employee）。
type Employee struct {
	ID         int64     `json:"id" db:"id"`
	Username   string    `json:"username" db:"username"`
	Name       string    `json:"name" db:"name"`
	Password   string    `json:"password" db:"password"`
	Phone      string    `json:"phone" db:"phone"`
	Sex        string    `json:"sex" db:"sex"`
	IDNumber   string    `json:"idNumber" db:"id_number"`
	Status     int       `json:"status" db:"status"`
	CreateTime time.Time `json:"createTime" db:"create_time"`
	UpdateTime time.Time `json:"updateTime" db:"update_time"`
	CreateUser *int64    `json:"createUser" db:"create_user"`
	UpdateUser *int64    `json:"updateUser" db:"update_user"`
}
