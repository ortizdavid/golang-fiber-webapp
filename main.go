package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-fiber-webapp/config"
	"github.com/ortizdavid/golang-fiber-webapp/controllers"
)

func main() {

	app := fiber.New(fiber.Config{
		Views: config.GetTemplateEngine(),
	})
	
	//Middlewares
	app.Use(controllers.RequestLoggerMiddleware)
	app.Use(controllers.AuthenticationMiddleware)

	config.ConfigStaticFiles(app)
	controllers.SetupRoutes(app)
	app.Listen(config.ListenAddr())
}