package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-fiber-webapp/config"
	"github.com/ortizdavid/golang-fiber-webapp/helpers"
	"github.com/ortizdavid/golang-fiber-webapp/models"
)

type ReportController struct {
}

var loggerReport = config.NewLogger("report-logs.log")

func (report ReportController) RegisterRoutes(router *fiber.App) {
	router.Get("/reports", report.reportHandler)
	router.Get("/reports/statistics", report.statisticsReportHandler)
}

func (ReportController) reportHandler(ctx *fiber.Ctx) error {
	param := ctx.Query("param")
	loggedUser := GetLoggedUser(ctx)
	var report models.ReportModel
	var tableReport models.TableReport
	var pdfGen helpers.HtmlPdfGenerator

	switch param {
	case "users":
		tableReport = report.GetAllUsers()
	case "tasks":
		tableReport = report.GetAllTasks()
	case "completed-tasks":
		tableReport = report.GetAllCompletedTasks()
	case "pending-tasks":
		tableReport = report.GetAllPendingTasks()
	case "in-progress-tasks":
		tableReport = report.GetAllInProgressTasks()
	case "cancelled-tasks":
		tableReport = report.GetAllCancelledTasks()
	case "blocked-tasks":
		tableReport = report.GetAllBlockedTasks()
	case "":
		return 	ctx.Render("reports/index", fiber.Map{
			"Title": "Reports",
			"LoggedUser": GetLoggedUser(ctx),
		})
	}
	//-----------------------
	templateFile :=  param +".html"
	title := "Report: " +tableReport.Title
	fileName := title +".pdf"
	data := map[string]any{
		"Title": title,
		"AppName": "Task Management App",
		"Rows": tableReport.Rows,
		"Count": tableReport.Count,
	}
	//----------- Render PDF
	pdfBytes, err := pdfGen.GeneratePDF(fmt.Sprintf("templates/reports/%s", templateFile), data)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	pdfGen.SetOutput(ctx, pdfBytes, fileName)
	loggerReport.Info(fmt.Sprintf("User '%s' generated '%s' Report", loggedUser.UserName, tableReport.Title))
	return nil
}

func (ReportController) statisticsReportHandler(ctx *fiber.Ctx) error {
	loggedUser := GetLoggedUser(ctx)
	var pdfGen helpers.HtmlPdfGenerator
	templateFile :=  "statistics.html"
	fileName := "Statistic Report.pdf"
	data := map[string]any{
		"Title": "Statistic Report",
		"AppName": "Task Management App",
		"Statistics": models.GetStatisticsCount(),
		"LoggedUser": GetLoggedUser(ctx),
	}
	//-----------------------
	pdfBytes, err := pdfGen.GeneratePDF(fmt.Sprintf("templates/reports/%s", templateFile), data)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	pdfGen.SetOutput(ctx, pdfBytes, fileName)
	loggerReport.Info(fmt.Sprintf("User '%s' generated '%s' Report", loggedUser.UserName, "Statistics"))
	return nil
}
