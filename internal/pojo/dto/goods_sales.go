package dto

// GoodsSalesDTO 商品销量数据传输对象（用于销量Top10统计）。
type GoodsSalesDTO struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}
