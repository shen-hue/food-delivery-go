package dto

import "time"

// DataOverViewQueryDTO 数据统计查询的数据传输对象。
type DataOverViewQueryDTO struct {
	Begin time.Time
	End   time.Time
}
