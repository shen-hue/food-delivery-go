package entity

// DishFlavor 菜品口味实体类（对应 Java 版 DishFlavor）。
// Value 为口味JSON字符串，如 '["微辣","中辣"]'。
type DishFlavor struct {
	ID     int64   `json:"id" db:"id"`
	DishID int64   `json:"dishId" db:"dish_id"`
	Name   *string `json:"name" db:"name"`
	Value  *string `json:"value" db:"value"`
}
