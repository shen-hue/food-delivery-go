package dto

import (
	"time"

	"sky-take-out/internal/pojo/entity"
)

// OrdersDTO 订单数据传输对象。
type OrdersDTO struct {
	ID            int64                `json:"id"`
	Number        string               `json:"number"`
	Status        int                  `json:"status"`
	UserID        int64                `json:"userId"`
	AddressBookID int64                `json:"addressBookId"`
	OrderTime     time.Time            `json:"orderTime"`
	CheckoutTime  *time.Time           `json:"checkoutTime"`
	PayMethod     *int                 `json:"payMethod"`
	Amount        float64              `json:"amount"`
	Remark        string               `json:"remark"`
	UserName      string               `json:"userName"`
	Phone         string               `json:"phone"`
	Address       string               `json:"address"`
	Consignee     string               `json:"consignee"`
	OrderDetails  []entity.OrderDetail `json:"orderDetails"`
}
