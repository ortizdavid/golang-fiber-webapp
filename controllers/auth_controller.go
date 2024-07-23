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

func (auth AuthController) RegisterRoutes(router *fiber.App) {
	group := router.Group("/auth")
	group.Get("/login", auth.loginForm)
	group.Post("/login", auth.login)
	group.Get("/logout", auth.logout)
	group.Get("/recover-password/:token", auth.recoverPasswordForm)
	group.Post("/recover-password/:token", auth.recoverPassword)
	group.Get("/get-recovery-link", auth.getRecoveryLinkForm)
	group.Post("/get-recovery-link", auth.getRecoveryLink)
}


func (AuthController) loginForm(c *fiber.Ctx) error {
	return c.Render("auth/login", fiber.Map{
		"Title": "Authentication",
	})
}

func (AuthController) login(c *fiber.Ctx) error {
	var userModel models.UserModel
	userName := c.FormValue("username")
	password := c.FormValue("password")
	user, _ := userModel.FindByUserName(userName)
	userExists, _ := userModel.ExistsActiveUser(userName)
	hashedPassword := user.Password
	
	if userExists && helpers.CheckPassword(hashedPassword, password) {
		store := config.GetSessionStore()
		session, _ := store.Get(c)
		session.Set("username", userName)
		session.Set("password", hashedPassword)
		session.Save()
		user.Token = helpers.GenerateRandomToken()
		userModel.Update(user)
		loggerAuth.Info(fmt.Sprintf("User '%s' authenticated sucessfully!", userName))
		return c.Redirect("/home")
	} else {
		loggerAuth.Error(fmt.Sprintf("User '%s' failed to login", userName))
		return c.Status(401).Redirect("/auth/login")
	}
}


func (AuthController) logout(c *fiber.Ctx) error {
	loggedUser := GetLoggedUser(c)
	store := config.GetSessionStore()
	session, err := store.Get(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	err = session.Destroy()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	loggerAuth.Info(fmt.Sprintf("User '%s' logged out", loggedUser.UserName))
	return c.Redirect("/auth/login")
}


func (AuthController) getRecoveryLinkForm(c *fiber.Ctx) error {
	return c.Render("auth/get-recovery-link", fiber.Map{
		"Title": "Recovery Link",
	})
}


func (AuthController) getRecoveryLink(c *fiber.Ctx) error {
	var userModel models.UserModel
	email := c.FormValue("email")
	user, err := userModel.FindByUserName(email)

	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	
	emailService := config.DefaultEmailService()
	recoverLink := fmt.Sprintf("http://%s/auth/recover-password/%s", config.ListenAddr(), user.Token)
	htmlBody := `
		<html>
			<body>
				<h1>Password recovery!</h1>
				<p>Hi, `+user.UserName+`!</p>
				<p>click on the link below <br> <a href="`+recoverLink+`">`+recoverLink+`</a></p>
			</body>
		</html>`
	err = emailService.SendHTMLEmail(email, "Password Recovery", htmlBody)
	if err != nil {
		fmt.Println("Error sending HTML email:", err)
	}
	loggerAuth.Info(fmt.Sprintf("User '%s' recovered password", email), config.LogRequestPath(c))
	return c.Redirect("/auth/get-recovery-link")
}


func (AuthController) recoverPasswordForm(c *fiber.Ctx) error {
	token := c.Params("token")
	user, err := models.UserModel{}.FindByToken(token)

	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.Render("auth/recover-password", fiber.Map{
		"Title": "Password Recovery",
		"User": user,
	})
}


func (AuthController) recoverPassword(c *fiber.Ctx) error {
	var userModel models.UserModel
	password := c.FormValue("password")
	token := c.Params("token")

	user, err := userModel.FindByToken(token)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	user.Password = helpers.HashPassword(password)
	user.Token = helpers.GenerateRandomToken()
	userModel.Update(user)

	//enviar os credenciais por email
	emailService := config.DefaultEmailService()
	htmlBody := `
		<html>
			<body>
				<h1>Password changed successfully!</h1>
				<p>Hi, `+user.UserName+`!</p>
				<p>New password is: <b>`+password+`</b></p>
			</body>
		</html>`
	err = emailService.SendHTMLEmail(user.UserName, "New Password", htmlBody)
	if err != nil {
		fmt.Println("Error sending HTML email:", err)
	}

	loggerAuth.Info(fmt.Sprintf("User '%s' recovered password", user.UserName), config.LogRequestPath(c))
	return c.Redirect("/auth/login")
}
