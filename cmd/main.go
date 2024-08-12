package main

import (
	"balancer/config"
	"balancer/internal/settings/server"
	"balancer/internal/settings/usecase"
	"context"
	"github.com/gofiber/fiber/v3"
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

	err = srv.Start(context.Background(), &config.C().ApiServer)
	if err != nil {
		panic(err)
	}

}
