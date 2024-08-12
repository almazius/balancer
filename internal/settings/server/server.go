package server

import (
	"balancer/config"
	"balancer/internal/settings/usecase"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"log/slog"
)

type SettingServer struct {
	app     *fiber.App
	service usecase.SettingService
}

func NewSettingServer(app *fiber.App, service usecase.SettingService) *SettingServer {
	return &SettingServer{app: app, service: service}
}

func (srv *SettingServer) Start(ctx context.Context, apiConfig *config.Server) error {
	var err error

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				err := srv.app.Shutdown()
				if err != nil {
					slog.Error("failed to shutdown api server", "error", err)
				}
			}
		}
	}(ctx)

	//err = srv.app.Listen(fmt.Sprintf("%s:%d", apiConfig.Host, apiConfig.Port))
	err = srv.app.Listen(fmt.Sprintf("%s:%d", "127.0.0.1", 8081))
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
