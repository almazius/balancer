package usecase

import (
	"balancer/internal/settings/usecase"
	"context"
	"fmt"
	"math/rand"
	"net/url"
	"strings"
)

type proxyService struct {
	proxySettings usecase.SettingService
}

func NewProxyService(proxySettings usecase.SettingService) ProxyService {
	return &proxyService{proxySettings: proxySettings}
}

func (s *proxyService) ProxyUrl(ctx context.Context, _url string) (string, error) {
	proxyMap, err := s.proxySettings.GetSetting(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get settings: %w", err)
	}

	u, err := url.Parse(_url)
	if err != nil {
		return "", fmt.Errorf("failed to parse url: %w", err)
	}

	path := u.Path + "/"

	for strings.LastIndex(path, "/") != -1 {
		path = path[:strings.LastIndex(path, "/")]
		proxyPaths, exists := proxyMap.Proxies[path]
		if exists {
			return proxyPaths[rand.Int()%len(proxyPaths)], nil // Пока только рандом
		}
	}

	return _url, nil
}
