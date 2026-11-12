package entity

import "time"

// 订单状态常量（对应 Java 版 Orders 中的静态常量）。
const (
	OrderStatusPendingPayment     = 1 // 待付款
	OrderStatusToBeConfirmed      = 2 // 待接单
	OrderStatusConfirmed          = 3 // 已接单
	OrderStatusDeliveryInProgress = 4 // 派送中
	OrderStatusCompleted          = 5 // 已完成
	OrderStatusCancelled          = 6 // 已取消
)

// 支付状态常量。
const (
	PayStatusUnPaid = 0 // 未支付
	PayStatusPaid   = 1 // 已支付
	PayStatusRefund = 2 // 退款
)

// Orders 订单实体类（对应 Java 版 Orders）。
type Orders struct {
	ID                    int64      `json:"id" db:"id"`
	Number                string     `json:"number" db:"number"`
	Status                int        `json:"status" db:"status"`
	UserID                int64      `json:"userId" db:"user_id"`
	AddressBookID         int64      `json:"addressBookId" db:"address_book_id"`
	OrderTime             time.Time  `json:"orderTime" db:"order_time"`
	CheckoutTime          *time.Time `json:"checkoutTime" db:"checkout_time"`
	PayMethod             *int       `json:"payMethod" db:"pay_method"` // 1微信 2支付宝
	PayStatus             int        `json:"payStatus" db:"pay_status"`
	Amount                float64    `json:"amount" db:"amount"`
	Remark                *string    `json:"remark" db:"remark"`
	UserName              string     `json:"userName" db:"user_name"`
	Phone                 string     `json:"phone" db:"phone"`
	Address               string     `json:"address" db:"address"`
	Consignee             string     `json:"consignee" db:"consignee"`
	CancelReason          *string    `json:"cancelReason" db:"cancel_reason"`
	RejectionReason       *string    `json:"rejectionReason" db:"rejection_reason"`
	CancelTime            *time.Time `json:"cancelTime" db:"cancel_time"`
	EstimatedDeliveryTime *time.Time `json:"estimatedDeliveryTime" db:"estimated_delivery_time"`
	DeliveryStatus        *int       `json:"deliveryStatus" db:"delivery_status"` // 1立即送出 0选择具体时间
	DeliveryTime          *time.Time `json:"deliveryTime" db:"delivery_time"`
	PackAmount            int        `json:"packAmount" db:"pack_amount"`           // 打包费
	TablewareNumber       int        `json:"tablewareNumber" db:"tableware_number"` // 餐具数量
	TablewareStatus       *int       `json:"tablewareStatus" db:"tableware_status"` // 1按餐量提供 0选择具体数量
}
