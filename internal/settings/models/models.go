package models

type ProxyDTO struct {
	InputPath string   `json:"inputPath"`
	ProxyUrls []string `json:"proxyUrls"`
}

type DeleteProxyDTO struct {
	InputPath string `json:"inputPath"`
	ProxyUrl  string `json:"proxyUrl"`
}
