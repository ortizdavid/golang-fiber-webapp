package controllers

import "github.com/gofiber/fiber/v2"

type FrontController struct {
}

func (FrontController) index(ctx *fiber.Ctx) error {
	return ctx.Redirect("/auth/login")
}

func (front FrontController) RegisterRoutes(router *fiber.App) {
	router.Get("/", front.index)
}