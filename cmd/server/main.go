// Package main 对应 Java 版 sky-server 的 SkyApplication（应用启动类）。
//
// 启动方式（在项目根目录下执行）：
//
//	go run ./cmd/server
package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"sky-take-out/internal/common/properties"
	"sky-take-out/internal/server/config"
	"sky-take-out/internal/server/controller/admin"
	"sky-take-out/internal/server/interceptor"
	"sky-take-out/internal/server/mapper"
	"sky-take-out/internal/server/service/impl"
)

func main() {
	// 加载配置（对应 Spring Boot 读取 application.yml + profile）
	profile := properties.ProfileFromEnv()
	cfg, err := properties.Load(profile)
	if err != nil {
		slog.Error("加载配置失败", "err", err)
		os.Exit(1)
	}

	// 初始化数据库连接池（对应 Spring 的 DataSource）
	pool, err := pgxpool.New(context.Background(), cfg.DSN())
	if err != nil {
		slog.Error("初始化数据库连接池失败", "err", err)
		os.Exit(1)
	}
	defer pool.Close()

	// 组装依赖（对应 Spring 的依赖注入容器）
	employeeMapper := mapper.NewEmployeeMapper(pool)
	employeeService := impl.NewEmployeeServiceImpl(employeeMapper)
	employeeController := admin.NewEmployeeController(employeeService, &cfg.Sky.Jwt)
	jwtTokenAdminInterceptor := interceptor.NewJwtTokenAdminInterceptor(&cfg.Sky.Jwt)
	router := config.NewRouter(employeeController, jwtTokenAdminInterceptor)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}

	// 启动http服务
	go func() {
		slog.Info("server started", "port", cfg.Server.Port, "profile", profile)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server启动失败", "err", err)
			os.Exit(1)
		}
	}()

	// 优雅停机（收到退出信号后等待请求处理完成再关闭）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("server关闭失败", "err", err)
	}
	slog.Info("server exited")
}
