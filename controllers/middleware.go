package controllers

import (
	"fmt"
	"strings"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-fiber-webapp/config"
	"go.uber.org/zap"
)


var requestLogger = config.NewLogger("requests.log")

func RequestLoggerMiddleware(ctx *fiber.Ctx) error {
	requestLogger.Info("Request",
		zap.String("Method", ctx.Method()),
		zap.String("Path", ctx.Path()),
		zap.String("StatusCode", fmt.Sprintf("%d", ctx.Response().StatusCode())),
	)
	return ctx.Next()
}

func AuthenticationMiddleware(ctx *fiber.Ctx) error {
	requestedPath := ctx.Path()
	if requestedPath ==  "/" || 
		strings.HasPrefix(requestedPath, "/image") ||
		strings.HasPrefix(requestedPath, "/css") ||
		strings.HasPrefix(requestedPath, "/js") ||
		strings.HasPrefix(requestedPath, "/auth") {
		return ctx.Next()
	}
	if !IsUserAuthenticated(ctx) {
		loggerAuth.Error(fmt.Sprintf("Authentication failed at: %s", requestedPath))
		return ctx.Status(500).Render("error/authentication", fiber.Map{
			"Title": "Authentication Error",
		})
	}
	return ctx.Next()
}


func IsUserAuthenticated(ctx *fiber.Ctx) bool {
	loggedUser := GetLoggedUser(ctx)
	if loggedUser.UserId == 0 && loggedUser.RoleId == 0 {
		return false
	}
	return true
}
