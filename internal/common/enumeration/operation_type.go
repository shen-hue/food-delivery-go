// Package enumeration 对应 Java 版的 com.sky.enumeration 包。
package enumeration

// OperationType 数据库操作类型（对应 Java 版的 OperationType 枚举）。
//
// 说明：Java 版配合自定义注解 @AutoFill + AOP切面 AutoFillAspect，
// 在 insert/update 时自动填充公共字段（createTime、createUser等）。
// Go 中没有 AOP，练习时可以直接在 mapper/service 中设置这些字段，
// 或者封装一个类似 autofill.Employee(emp, op) 的小工具函数。
type OperationType int

const (
	OperationTypeUpdate OperationType = iota // UPDATE 更新操作
	OperationTypeInsert                      // INSERT 插入操作
)
