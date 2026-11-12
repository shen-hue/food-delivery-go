// Package basecontext 对应 Java 版的 com.sky.context.BaseContext。
//
// Java 版使用 ThreadLocal 在当前线程中保存登录用户id；
// Go 中没有 ThreadLocal，对应的惯用法是把值放进 context.Context，
// 随请求一路传递（request-scoped）。
//
// 用法：
//
//	ctx = basecontext.WithCurrentID(ctx, empID) // 拦截器/中间件中写入
//	empID := basecontext.CurrentID(ctx)          // service/mapper中读取
package basecontext

import "context"

type contextKey struct{}

var currentIDKey contextKey

// WithCurrentID 对应 BaseContext.setCurrentId —— 返回携带当前登录用户id的新context。
func WithCurrentID(ctx context.Context, id int64) context.Context {
	return context.WithValue(ctx, currentIDKey, id)
}

// CurrentID 对应 BaseContext.getCurrentId —— 读取当前登录用户id，未设置时返回0。
func CurrentID(ctx context.Context) int64 {
	if v, ok := ctx.Value(currentIDKey).(int64); ok {
		return v
	}
	return 0
}
