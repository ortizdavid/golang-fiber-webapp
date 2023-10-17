package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-fiber-webapp/models"
)

type BackController struct {
}

func (back BackController) RegisterRoutes(router *fiber.App) {
	router.Get("/home", back.home)
	router.Get("/user-data", back.userData)
	router.Get("/change-password", back.changePassword)
	router.Get("/upload-image", back.uploadImage)
}

func (BackController) home(ctx *fiber.Ctx) error {
	loggedUser := GetLoggedUser(ctx)
	statistics := models.GetStatisticsCount()
	if IsUserNomal(ctx) {
		statistics = models.GetStatisticsCountByUser(loggedUser.UserId)
	}
	return ctx.Render("back-office/home", fiber.Map{
		"Title": "Home",
		"Statistics": statistics,
		"LoggedUser": loggedUser,
	})
}

func (BackController) userData(ctx *fiber.Ctx) error {
	return ctx.Render("back-office/user-data", fiber.Map{
		"Title": "User Data",
		"LoggedUser": GetLoggedUser(ctx),
	})
}

func (BackController) changePassword(ctx *fiber.Ctx) error {
	return ctx.Render("back-office/change-password", fiber.Map{
		"Title": "Change Password",
		"LoggedUser": GetLoggedUser(ctx),
	})
}

func (BackController) uploadImage(ctx *fiber.Ctx) error {
	return ctx.Render("back-office/upload-image", fiber.Map{
		"Title": "Upload Image",
		"LoggedUser": GetLoggedUser(ctx),
	})
}