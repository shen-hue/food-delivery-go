package entity

// SetmealDish 套餐菜品关系实体类（对应 Java 版 SetmealDish）。
type SetmealDish struct {
	ID        int64    `json:"id" db:"id"`
	SetmealID int64    `json:"setmealId" db:"setmeal_id"`
	DishID    int64    `json:"dishId" db:"dish_id"`
	Name      *string  `json:"name" db:"name"`
	Price     *float64 `json:"price" db:"price"`
	Copies    *int     `json:"copies" db:"copies"`
}
