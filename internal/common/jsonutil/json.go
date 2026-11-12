// Package jsonutil 对应 Java 版中 Jackson（JacksonObjectMapper）承担的职责：
// 请求体解析与响应体序列化。
//
// 与 Java 版一致的行为：
//   - 解析请求体时忽略未知字段（Jackson 默认 FAIL_ON_UNKNOWN_PROPERTIES=false）
//   - 响应统一输出 application/json
package jsonutil

import (
	"encoding/json"
	"net/http"
)

// WriteJSON 向响应写入JSON数据（相当于 @ResponseBody + Jackson序列化）。
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// ReadJSON 解析请求体（相当于 @RequestBody 反序列化）。
func ReadJSON(r *http.Request, v any) error {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(v)
}
