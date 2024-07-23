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

func (BackController) home(c *fiber.Ctx) error {
	loggedUser := GetLoggedUser(c)
	statistics := models.GetStatisticsCount()
	if IsUserNomal(c) {
		statistics = models.GetStatisticsCountByUser(loggedUser.UserId)
	}
	return c.Render("back-office/home", fiber.Map{
		"Title": "Home",
		"Statistics": statistics,
		"LoggedUser": loggedUser,
	})
}

func (BackController) userData(c *fiber.Ctx) error {
	return c.Render("back-office/user-data", fiber.Map{
		"Title": "User Data",
		"LoggedUser": GetLoggedUser(c),
	})
}

func (BackController) changePassword(c *fiber.Ctx) error {
	return c.Render("back-office/change-password", fiber.Map{
		"Title": "Change Password",
		"LoggedUser": GetLoggedUser(c),
	})
}

func (BackController) uploadImage(c *fiber.Ctx) error {
	return c.Render("back-office/upload-image", fiber.Map{
		"Title": "Upload Image",
		"LoggedUser": GetLoggedUser(c),
	})
}