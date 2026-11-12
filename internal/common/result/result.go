// Package result 对应 Java 版 sky-common 的 com.sky.result 包。
package result

// Result 后端统一返回结果（对应 Java 版的 Result<T>）。
//
// Go 版说明：Java 泛型在运行时被擦除，这里直接使用 any 承载数据，
// JSON 输出与 Java 版完全一致：{"code":1,"msg":null,"data":{...}}。
//
// 编码：1成功，0和其它数字为失败。
type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// Success 对应 Result.success(object) —— 成功并返回数据。
func Success(data any) Result {
	return Result{Code: 1, Data: data}
}

// SuccessEmpty 对应 Result.success() —— 成功但不返回数据。
func SuccessEmpty() Result {
	return Result{Code: 1}
}

// Error 对应 Result.error(msg) —— 失败并返回错误信息。
func Error(msg string) Result {
	return Result{Code: 0, Msg: msg}
}
