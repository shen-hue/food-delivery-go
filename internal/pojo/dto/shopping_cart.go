package dto

// ShoppingCartDTO 添加购物车的数据传输对象。
type ShoppingCartDTO struct {
	DishID     *int64 `json:"dishId"`
	SetmealID  *int64 `json:"setmealId"`
	DishFlavor string `json:"dishFlavor"`
}
