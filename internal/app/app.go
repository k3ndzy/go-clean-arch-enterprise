package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/k3ndzy/go-clean-arch-enterprise/internal/config"
	deliveryHTTP "github.com/k3ndzy/go-clean-arch-enterprise/internal/delivery/http"
	"github.com/k3ndzy/go-clean-arch-enterprise/internal/delivery/http/handler"
	"github.com/k3ndzy/go-clean-arch-enterprise/internal/delivery/http/middleware"
	"github.com/k3ndzy/go-clean-arch-enterprise/pkg/logger"
)

// Run khởi động toàn bộ ứng dụng và ráp nối Dependency Injection (Plan 004)
func Run() error {
	// 1. Tải cấu hình
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("không thể tải cấu hình: %w", err)
	}

	// 2. Khởi tạo Structured Logger
	logLevel := slog.LevelInfo
	if cfg.App.Environment == config.EnvDevelopment {
		logLevel = slog.LevelDebug
	}
	appLogger := logger.New(os.Stdout, logLevel)

	appLogger.Info("khởi động ứng dụng",
		slog.String("app", cfg.App.Name),
		slog.String("version", cfg.App.Version),
		slog.String("env", string(cfg.App.Environment)),
	)

	// 3. Khởi tạo HTTP Router & Handlers
	mux := http.NewServeMux()

	healthHandler := handler.NewHealthHandler(cfg.App.Version)
	healthHandler.RegisterRoutes(mux)

	// 4. Áp dụng Middlewares (Request ID -> Logging -> Routing)
	var rootHandler http.Handler = mux
	rootHandler = middleware.Logger(appLogger)(rootHandler)
	rootHandler = middleware.RequestID(rootHandler)

	// 5. Khởi tạo & Chạy HTTP Server với Graceful Shutdown
	srv := deliveryHTTP.NewServer(cfg, rootHandler, appLogger)
	return srv.Start()
}
