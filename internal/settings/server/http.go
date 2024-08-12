package server

import "github.com/gofiber/fiber/v3"

func MapRoutes(s *SettingServer, app *fiber.App) {
	apiRoute := app.Group("/api")
	settingsRoute := apiRoute.Group("/settings") // todo mw for checking auth

	settingsRoute.Get("/", s.GetSettings()) // return settings

	settingsRoute.Post("/proxy", s.AddProxy())                 // add new proxy
	settingsRoute.Delete("/proxy", s.DeleteProxyPath())        // del new proxy
	settingsRoute.Delete("/input", s.DeleteProxyByInputPath()) // del new input on proxy path

}
