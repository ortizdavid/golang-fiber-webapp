package controllers

import (
	"fmt"
	"strings"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-fiber-webapp/config"
	"go.uber.org/zap"
)


var requestLogger = config.NewLogger("requests.log")

func RequestLoggerMiddleware(c *fiber.Ctx) error {
	requestLogger.Info("Request",
		zap.String("Method", c.Method()),
		zap.String("Path", c.Path()),
		zap.String("StatusCode", fmt.Sprintf("%d", c.Response().StatusCode())),
	)
	return c.Next()
}

func AuthenticationMiddleware(c *fiber.Ctx) error {
	requestedPath := c.Path()
	if requestedPath ==  "/" || 
		strings.HasPrefix(requestedPath, "/image") ||
		strings.HasPrefix(requestedPath, "/css") ||
		strings.HasPrefix(requestedPath, "/js") ||
		strings.HasPrefix(requestedPath, "/auth") {
		return c.Next()
	}
	if !IsUserAuthenticated(c) {
		loggerAuth.Error(fmt.Sprintf("Authentication failed at: %s", requestedPath))
		return c.Status(500).Render("error/authentication", fiber.Map{
			"Title": "Authentication Error",
		})
	}
	return c.Next()
}


func IsUserAuthenticated(c *fiber.Ctx) bool {
	loggedUser := GetLoggedUser(c)
	if loggedUser.UserId == 0 && loggedUser.RoleId == 0 {
		return false
	}
	return true
}
