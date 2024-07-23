package controllers

import "github.com/gofiber/fiber/v2"

type InterfaceController interface {
	index(c *fiber.Ctx) error
	details(c *fiber.Ctx) error
	addForm(c *fiber.Ctx) error
	add(c *fiber.Ctx) error
	editForm(c *fiber.Ctx) error
	edit(c *fiber.Ctx) error
	searchForm(c *fiber.Ctx) error
	search(c *fiber.Ctx) error
	RegisterRoutes(router *fiber.App)
}