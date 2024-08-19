package main

import (
	"balancer/config"
	"balancer/internal/proxy/handler"
	usecase2 "balancer/internal/proxy/usecase"
	"balancer/internal/settings/server"
	"balancer/internal/settings/usecase"
	"context"
	"github.com/gofiber/fiber/v3"
	"sync"
)

func main() {
	_, err := config.ParseConfig()
	if err != nil {
		panic(err)
	}

	service := usecase.NewSettingService()
	app := fiber.New()
	srv := server.NewSettingServer(app, service)
	server.MapRoutes(srv, app)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		err = srv.Start(context.Background(), &config.C().ApiServer)
		if err != nil {
			panic(err)
		}
		wg.Done()
	}(wg)

	proxyApp := fiber.New()
	proxyService := usecase2.NewProxyService(service)
	proxyServer := handler.NewProxyServer(proxyApp, proxyService)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		err = proxyServer.Start(context.Background(), &config.C().ApiServer)
		if err != nil {
			panic(err)
		}
		wg.Done()
	}(wg)

	wg.Wait()
}
