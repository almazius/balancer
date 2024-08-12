package server

import (
	"balancer/internal/settings/models"
	"balancer/pkg/utils/bodyparser"
	"github.com/gofiber/fiber/v3"
	"log/slog"
)

func (srv *SettingServer) GetSettings() fiber.Handler {
	return func(c fiber.Ctx) error {
		settings, err := srv.service.GetSetting(c.Context())
		if err != nil {
			slog.Error("failed getting settings", "error", err)
			return err
		}

		return c.JSON(settings)
	}
}

func (srv *SettingServer) AddProxy() fiber.Handler {
	return func(c fiber.Ctx) error {
		proxyInfo := new(models.ProxyDTO)
		err := bodyparser.ParseBody(c, proxyInfo)
		if err != nil {
			slog.Error("failed parse body", "error", err)
			return err
		}

		err = srv.service.AddProxy(c.Context(), proxyInfo.InputPath, proxyInfo.ProxyUrls)
		if err != nil {
			slog.Error("failed adding proxy", "error", err)
			return err
		}

		return c.JSON(fiber.Map{"message": "proxy added"})
	}
}

func (srv *SettingServer) DeleteProxyByInputPath() fiber.Handler {
	return func(c fiber.Ctx) error {
		proxyInfo := new(models.DeleteProxyDTO)
		err := bodyparser.ParseBody(c, proxyInfo)
		if err != nil {
			slog.Error("failed parse body", "error", err)
			return err
		}

		err = srv.service.DeleteInputURL(c.Context(), proxyInfo.InputPath)
		if err != nil {
			slog.Error("failed adding proxy", "error", err)
			return err
		}

		return c.JSON(fiber.Map{"message": "proxy added"})
	}
}

func (srv *SettingServer) DeleteProxyPath() fiber.Handler {
	return func(c fiber.Ctx) error {
		proxyInfo := new(models.DeleteProxyDTO)
		err := bodyparser.ParseBody(c, proxyInfo)
		if err != nil {
			slog.Error("failed parse body", "error", err)
			return err
		}

		err = srv.service.DeleteProxyURL(c.Context(), proxyInfo.InputPath, proxyInfo.ProxyUrl)
		if err != nil {
			slog.Error("failed adding proxy", "error", err)
			return err
		}

		return c.JSON(fiber.Map{"message": "proxy added"})
	}
}
