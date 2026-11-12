package dto

import "sky-take-out/internal/common/jsonutil"

// OrdersSubmitDTO 用户下单的数据传输对象。
type OrdersSubmitDTO struct {
	AddressBookID         int64             `json:"addressBookId"`
	PayMethod             int               `json:"payMethod"`
	Remark                string            `json:"remark"`
	EstimatedDeliveryTime jsonutil.DateTime `json:"estimatedDeliveryTime"` // yyyy-MM-dd HH:mm:ss
	DeliveryStatus        *int              `json:"deliveryStatus"`        // 1立即送出 0选择具体时间
	TablewareNumber       *int              `json:"tablewareNumber"`       // 餐具数量
	TablewareStatus       *int              `json:"tablewareStatus"`       // 1按餐量提供 0选择具体数量
	PackAmount            int               `json:"packAmount"`            // 打包费
	Amount                float64           `json:"amount"`                // 总金额
}
