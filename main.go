package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-fiber-webapp/config"
	//"github.com/ortizdavid/golang-fiber-webapp/entities"
	"github.com/ortizdavid/golang-fiber-webapp/controllers"
)

func main() {

	app := fiber.New(fiber.Config{
		Views: config.GetTemplateEngine(),
	})
	//entities.SetupMigrations()
	config.ConfigStaticFiles(app)
	controllers.SetupRoutes(app)
	app.Listen(config.ListenAddr())
}