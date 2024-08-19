package usecase

import "context"

type ProxyService interface {
	ProxyUrl(ctx context.Context, url string) (string, error)
}
