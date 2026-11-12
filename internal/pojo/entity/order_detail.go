package entity

// OrderDetail 订单明细实体类（对应 Java 版 OrderDetail）。
type OrderDetail struct {
	ID         int64   `json:"id" db:"id"`
	Name       *string `json:"name" db:"name"`
	OrderID    int64   `json:"orderId" db:"order_id"`
	DishID     *int64  `json:"dishId" db:"dish_id"`
	SetmealID  *int64  `json:"setmealId" db:"setmeal_id"`
	DishFlavor *string `json:"dishFlavor" db:"dish_flavor"`
	Number     int     `json:"number" db:"number"`
	Amount     float64 `json:"amount" db:"amount"`
	Image      *string `json:"image" db:"image"`
}
