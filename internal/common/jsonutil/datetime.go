package jsonutil

import (
	"fmt"
	"time"
)

// DateTimeLayout 对应 Java DTO 上的 @JsonFormat(pattern = "yyyy-MM-dd HH:mm:ss")
// 与 @DateTimeFormat(pattern = "yyyy-MM-dd HH:mm:ss") 注解的格式。
const DateTimeLayout = "2006-01-02 15:04:05"

// DateTime 包装 time.Time，使其按 "yyyy-MM-dd HH:mm:ss" 格式进行JSON序列化/反序列化。
//
// Go 标准库的 time.Time 默认按 RFC3339（含时区）序列化，
// 而前端传入/期望的是 "2026-11-10 12:30:00" 这种格式，
// 因此在 DTO/VO 中需要自定义格式的字段使用本类型：
//
//	EstimatedDeliveryTime jsonutil.DateTime `json:"estimatedDeliveryTime"`
type DateTime struct {
	time.Time
}

// MarshalJSON 实现 json.Marshaler。
func (t DateTime) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + t.Time.Format(DateTimeLayout) + `"`), nil
}

// UnmarshalJSON 实现 json.Unmarshaler。
func (t *DateTime) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "null" || s == `""` {
		t.Time = time.Time{}
		return nil
	}
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}
	parsed, err := time.ParseInLocation(DateTimeLayout, s, time.Local)
	if err != nil {
		return fmt.Errorf("datetime格式应为 %q: %w", DateTimeLayout, err)
	}
	t.Time = parsed
	return nil
}
