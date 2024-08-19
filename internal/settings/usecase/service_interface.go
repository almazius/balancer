package usecase

import (
	"balancer/internal/settings/models"
	"context"
)

type SettingService interface {
	GetSetting(ctx context.Context) (*models.ProxySettingsDTO, error)

	AddProxy(ctx context.Context, inputUrl string, proxyUrl []string) error
	DeleteInputURL(ctx context.Context, inputURL string) error
	DeleteProxyURL(ctx context.Context, inputURL, proxyURL string) error
}
