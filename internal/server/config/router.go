// Package config 对应 Java 版 sky-server 的 com.sky.config 包
// （WebMvcConfiguration：注册拦截器、路由等web层组件）。
package config

import (
	"log/slog"
	"net/http"
	"time"

	"sky-take-out/internal/server/controller/admin"
	"sky-take-out/internal/server/handler"
	"sky-take-out/internal/server/interceptor"
)

// NewRouter 构建路由与中间件
// （对应 Java 版 WebMvcConfiguration.addInterceptors 中的注册逻辑：
// 拦截 /admin/** 路径，放行 /admin/employee/login）。
func NewRouter(
	employeeController *admin.EmployeeController,
	jwtTokenAdminInterceptor *interceptor.JwtTokenAdminInterceptor,
) http.Handler {
	mux := http.NewServeMux()

	// 放行的接口（对应 excludePathPatterns("/admin/employee/login")）
	// Go 1.22 的路由规则：精确的 "POST /admin/employee/login"
	// 优先级高于 "/admin/" 前缀匹配，因此不会经过jwt中间件
	mux.Handle("POST /admin/employee/login", handler.Wrap(employeeController.Login))

	// 管理端接口，统一经过jwt令牌校验中间件
	adminMux := http.NewServeMux()
	adminMux.Handle("POST /admin/employee/logout", handler.Wrap(employeeController.Logout))
	// TODO: 后续练习中在此注册其它管理端接口（/admin/...）
	mux.Handle("/admin/", jwtTokenAdminInterceptor.Middleware(adminMux))

	// TODO: 用户端接口（/user/...）需要另一个 JwtTokenUserInterceptor，自己动手实现吧

	// 请求日志中间件（对应 Spring 的 request logging，方便调试）
	return logMiddleware(mux)
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		slog.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"duration", time.Since(start).String())
	})
}
