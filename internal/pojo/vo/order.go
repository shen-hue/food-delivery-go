package vo

import "sky-take-out/internal/pojo/entity"

// OrderVO 订单展示对象。
//
// 对应 Java 版的 OrderVO extends Orders：
// Go 中通过嵌入 entity.Orders 实现同样的字段继承效果，
// JSON序列化时订单字段与新增字段平铺输出，与 Java 版一致。
type OrderVO struct {
	entity.Orders
	OrderDishes     string               `json:"orderDishes"`     // 订单菜品信息（字符串形式）
	OrderDetailList []entity.OrderDetail `json:"orderDetailList"` // 订单明细列表
}
