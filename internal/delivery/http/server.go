package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/k3ndzy/go-clean-arch-enterprise/internal/config"
)

type Server struct {
	httpServer *http.Server
	cfg        *config.Config
	logger     *slog.Logger
}

func NewServer(cfg *config.Config, handler http.Handler, log *slog.Logger) *Server {
	return &Server{
		cfg:    cfg,
		logger: log,
		httpServer: &http.Server{
			Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
			Handler:      handler,
			ReadTimeout:  cfg.Server.ReadTimeout,
			WriteTimeout: cfg.Server.WriteTimeout,
			IdleTimeout:  cfg.Server.IdleTimeout,
		},
	}
}

// Start khởi chạy HTTP server và lắng nghe tín hiệu tắt nhẹ nhàng (Graceful Shutdown)
func (s *Server) Start() error {
	shutdownErr := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		sig := <-quit

		s.logger.Info("nhận tín hiệu dừng server, bắt đầu đóng các kết nối...", slog.String("signal", sig.String()))

		ctx, cancel := context.WithTimeout(context.Background(), s.cfg.Server.ShutdownTimeout)
		defer cancel()

		shutdownErr <- s.httpServer.Shutdown(ctx)
	}()

	s.logger.Info("🚀 HTTP Server đang lắng nghe",
		slog.Int("port", s.cfg.Server.Port),
		slog.String("env", string(s.cfg.App.Environment)),
	)

	err := s.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	err = <-shutdownErr
	if err != nil {
		return fmt.Errorf("lỗi trong quá trình graceful shutdown: %w", err)
	}

	s.logger.Info("✔ Server đã dừng hoàn toàn an toàn.")
	return nil
}
