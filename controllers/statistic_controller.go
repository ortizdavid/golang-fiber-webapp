package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-fiber-webapp/models"
)

type StatisticController struct {
}

func (statistic StatisticController) RegisterRoutes(router *fiber.App) {
	router.Get("/statistics", statistic.index)
}

func (StatisticController) index(ctx *fiber.Ctx) error {
	loggedUser := GetLoggedUser(ctx)
	statistics :=  models.GetStatisticsCount()
	if loggedUser.RoleCode == "normal" {
		statistics = models.GetStatisticsCountByUser(loggedUser.UserId)
	}
	return ctx.Render("statistic/index", fiber.Map{
		"Title": "Statistics",
		"LoggedUser": loggedUser,
		"Statistics": statistics,
	})
}