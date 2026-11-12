// Package interceptor 对应 Java 版 sky-server 的 com.sky.interceptor 包。
package interceptor

import (
	"log/slog"
	"net/http"

	"sky-take-out/internal/common/constant"
	"sky-take-out/internal/common/properties"
	"sky-take-out/internal/common/utils"
)

// JwtTokenAdminInterceptor 管理端jwt令牌校验中间件
// （对应 Java 版 JwtTokenAdminInterceptor 实现 HandlerInterceptor）。
//
// Go 中通过中间件函数包裹 http.Handler 实现拦截器效果。
type JwtTokenAdminInterceptor struct {
	jwtProperties *properties.JwtProperties
}

// NewJwtTokenAdminInterceptor 构造函数（对应 @Component + @Autowired 注入）。
func NewJwtTokenAdminInterceptor(jwtProperties *properties.JwtProperties) *JwtTokenAdminInterceptor {
	return &JwtTokenAdminInterceptor{jwtProperties: jwtProperties}
}

// Middleware 返回管理端jwt校验中间件。
//
// 对应 Java 版 preHandle 的逻辑：
//  1. 从请求头中获取令牌
//  2. 校验令牌，通过则放行，不通过则响应401状态码
func (i *JwtTokenAdminInterceptor) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1、从请求头中获取令牌
		token := r.Header.Get(i.jwtProperties.AdminTokenName)

		// 2、校验令牌
		claims, err := utils.ParseJWT(i.jwtProperties.AdminSecretKey, token)
		if err != nil {
			// 4、不通过，响应401状态码
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		slog.Info("jwt校验通过", "empId", claims[constant.JwtClaimEmpID])

		// 3、通过，放行
		// TODO: 与Java完整版一样，这里可以把当前员工id存入请求context，
		//       供service/mapper层使用（公共字段自动填充时需要）：
		//       r = r.WithContext(basecontext.WithCurrentID(r.Context(), empID))
		next.ServeHTTP(w, r)
	})
}
