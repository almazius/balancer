package models

type ProxySettingsDTO struct {
	Proxies map[string][]string `json:"proxies"`
}
