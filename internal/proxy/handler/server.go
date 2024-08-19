package handler

import (
	"balancer/config"
	"balancer/internal/proxy/usecase"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"log/slog"
)

var AllMethods = []string{
	fiber.MethodGet,
	fiber.MethodHead,
	fiber.MethodPost,
	fiber.MethodPut,
	fiber.MethodPatch,
	fiber.MethodDelete,
	fiber.MethodConnect,
	fiber.MethodOptions,
	fiber.MethodTrace,
}

type ProxyServer struct {
	app     *fiber.App
	service usecase.ProxyService
}

func NewProxyServer(app *fiber.App, service usecase.ProxyService) *ProxyServer {
	return &ProxyServer{app: app, service: service}
}

func (srv *ProxyServer) Start(ctx context.Context, apiConfig *config.Server) error {
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

	//srv.app.Add(AllMethods, "/", srv.Proxy())
	srv.app.Use("/", srv.Proxy())

	//err = srv.app.Listen(fmt.Sprintf("%s:%d", apiConfig.Host, apiConfig.Port))
	err = srv.app.Listen(fmt.Sprintf("%s:%d", "127.0.0.1", 8080))
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

func (srv *ProxyServer) Proxy() fiber.Handler {
	return func(c fiber.Ctx) error {
		Url := c.OriginalURL()

		resultUrl, err := srv.service.ProxyUrl(c.Context(), Url)
		if err != nil {
			return err
		}

		slog.Info("test proxy", "first url", Url, "result url", resultUrl)

		return c.Redirect().To(resultUrl)

		//return c.JSON(resultUrl)
	}
}
