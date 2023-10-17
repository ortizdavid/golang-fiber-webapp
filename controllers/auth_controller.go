package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-fiber-webapp/config"
	"github.com/ortizdavid/golang-fiber-webapp/helpers"
	"github.com/ortizdavid/golang-fiber-webapp/models"
)

type AuthController struct {
}

var loggerAuth = config.NewLogger("auth-logs.log")

func (AuthController) loginForm(ctx *fiber.Ctx) error {
	return ctx.Render("auth/login", fiber.Map{
		"Title": "Authentication",
	})
}

func (AuthController) login(ctx *fiber.Ctx) error {
	var userModel models.UserModel
	userName := ctx.FormValue("username")
	password := ctx.FormValue("password")
	user := userModel.GetByUserName(userName)
	hashedPassword := user.Password
	
	if userModel.ExistsActiveUser(userName) && helpers.CheckPassword(hashedPassword, password) {
		store := config.GetSessionStore()
		session, _ := store.Get(ctx)
		session.Set("username", userName)
		session.Set("password", hashedPassword)
		session.Set("authenticated", true)
		session.Set("role", user.RoleName)
		session.Save()
		loggerAuth.Info(fmt.Sprintf("User '%s' authenticated sucessfully!", userName))
		return ctx.Redirect("/home")
	} else {
		loggerAuth.Error(fmt.Sprintf("User '%s' failed to login", userName))
		return ctx.Redirect("/auth/login")
	}
}

func (AuthController) logout(ctx *fiber.Ctx) error {
	loggedUser := GetLoggedUser(ctx)
	store := config.GetSessionStore()
	session, _ := store.Get(ctx)
	session.Destroy()
	loggerAuth.Info(fmt.Sprintf("User '%s' logged out", loggedUser.UserName))
	return ctx.Redirect("/auth/login")
}

func (AuthController) resetPassword(ctx *fiber.Ctx) error {
	return ctx.SendString("Processing Reset...")
}

func (AuthController) resetPasswordForm(ctx *fiber.Ctx) error {
	return ctx.Render("auth/reset-password", fiber.Map{
		"Title": "Reset Password",
	})
}

func (auth AuthController) RegisterRoutes(router *fiber.App) {
	group := router.Group("/auth")
	group.Get("/login", auth.loginForm)
	group.Post("/login", auth.login)
	group.Get("/logout", auth.logout)
	group.Get("/reset-password", auth.resetPasswordForm)
	group.Post("/reset-password", auth.resetPassword)
}

