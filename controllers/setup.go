package controllers

import "github.com/gofiber/fiber/v2"

func SetupRoutes(router *fiber.App) {
	FrontController{}.RegisterRoutes(router)
	AuthController{}.RegisterRoutes(router)
	BackController{}.RegisterRoutes(router)
	TaskController{}.RegisterRoutes(router)
	UserController{}.RegisterRoutes(router)
	ReportController{}.RegisterRoutes(router)
	StatisticController{}.RegisterRoutes(router)
}