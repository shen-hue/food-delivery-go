// Package handler 对应 Java 版 sky-server 的 com.sky.handler 包。
package handler

import (
	"errors"
	"log/slog"
	"net/http"

	"sky-take-out/internal/common/constant"
	"sky-take-out/internal/common/exception"
	"sky-take-out/internal/common/jsonutil"
	"sky-take-out/internal/common/result"
)

// HandlerFunc 业务处理函数签名。
//
// Go 的 net/http 处理函数默认不返回错误；
// 这里让Controller返回 error，由 Wrap 统一处理，
// 相当于 Java 中 Controller 抛出异常、由全局异常处理器捕获的机制。
type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

// Wrap 把 HandlerFunc 包装成标准 http.HandlerFunc。
//
// 对应 Java 版的 GlobalExceptionHandler（@RestControllerAdvice）：
// 捕获业务异常（*exception.BaseException 及其"子类"），
// 统一转换为 Result.error 响应；其它未知错误返回 500 + 未知错误提示。
func Wrap(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			var baseEx *exception.BaseException
			if errors.As(err, &baseEx) {
				// 业务异常：与 Java 版一致，HTTP 200 + code=0
				slog.Error("异常信息", "msg", baseEx.Msg)
				_ = jsonutil.WriteJSON(w, http.StatusOK, result.Error(baseEx.Msg))
				return
			}
			// 未知错误
			slog.Error("未知错误", "err", err)
			_ = jsonutil.WriteJSON(w, http.StatusInternalServerError, result.Error(constant.UnknownError))
		}
	}
}
