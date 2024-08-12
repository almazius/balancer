package models

type ProxyDTO struct {
	InputPath string
	ProxyUrls []string
}

type DeleteProxyDTO struct {
	InputPath string
	ProxyUrl  string
}
