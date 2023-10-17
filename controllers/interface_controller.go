package controllers

import "github.com/gofiber/fiber/v2"

type InterfaceController interface {
	index(ctx *fiber.Ctx) error
	details(ctx *fiber.Ctx) error
	addForm(ctx *fiber.Ctx) error
	add(ctx *fiber.Ctx) error
	editForm(ctx *fiber.Ctx) error
	edit(ctx *fiber.Ctx) error
	searchForm(ctx *fiber.Ctx) error
	search(ctx *fiber.Ctx) error
	RegisterRoutes(router *fiber.App)
}