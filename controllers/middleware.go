package controllers

import (
	"fmt"
	"strings"
	"github.com/gofiber/fiber/v2"
)


func AuthenticationMiddleware(ctx *fiber.Ctx) error {

	requestedPath := ctx.Path()
	if requestedPath ==  "/" || 
		strings.HasPrefix(requestedPath, "/images") ||
		strings.HasPrefix(requestedPath, "/css") ||
		strings.HasPrefix(requestedPath, "/js") ||
		strings.HasPrefix(requestedPath, "/auth") {
		return ctx.Next()
	}
	if !IsUserAuthenticated(ctx) {
		loggerAuth.Error(fmt.Sprintf("Authentication failed at: %s", requestedPath))
		return ctx.Render("error/authentication", fiber.Map{
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
