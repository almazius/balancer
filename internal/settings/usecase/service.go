package usecase

import (
	"balancer/pkg/consts"
	"context"
	"sync"
)

type settingService struct {
	mu       sync.RWMutex
	Settings *Settings
}

type Settings struct {
	balancerMap map[string][]string // key - input url, value - proxy url
}

func NewSettingService() SettingService {
	return &settingService{}
}

func (s *settingService) GetSetting(ctx context.Context) (*Settings, error) {
	return s.Settings, nil
}

func (s *settingService) AddProxy(ctx context.Context, inputUrl string, proxyUrl []string) error {
	s.mu.RLock()
	if len(s.Settings.balancerMap[inputUrl]) == 0 {
		s.Settings.balancerMap[inputUrl] = proxyUrl
	} else {
		s.Settings.balancerMap[inputUrl] = append(s.Settings.balancerMap[inputUrl], proxyUrl...)
	}

	s.mu.RUnlock()
	return nil
}

func (s *settingService) DeleteInputURL(ctx context.Context, inputURL string) error {
	s.mu.RLock()
	delete(s.Settings.balancerMap, inputURL)
	s.mu.RUnlock()
	return nil
}

func (s *settingService) DeleteProxyURL(ctx context.Context, inputURL, proxyURL string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tempProxyURL, exist := s.Settings.balancerMap[inputURL]
	if !exist {
		return consts.ErrNotFoundInputURL
	}

	delProxyIndex := -1
	for i := range tempProxyURL {
		if tempProxyURL[i] == proxyURL {
			delProxyIndex = i
			break
		}
	}

	if delProxyIndex == -1 {
		return consts.ErrNotFoundProxyURL
	}

	if delProxyIndex == 0 {
		tempProxyURL = tempProxyURL[1:]
	} else if delProxyIndex == len(tempProxyURL)-1 {
		tempProxyURL = tempProxyURL[:delProxyIndex]
	} else {
		tempProxyURL = append(tempProxyURL[:delProxyIndex], tempProxyURL[delProxyIndex+1:]...)
	}

	s.Settings.balancerMap[inputURL] = tempProxyURL

	if len(s.Settings.balancerMap[inputURL]) == 0 {
		delete(s.Settings.balancerMap, inputURL)
	}
	return nil
}
