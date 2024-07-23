package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-fiber-webapp/config"
	"github.com/ortizdavid/golang-fiber-webapp/entities"
	"github.com/ortizdavid/golang-fiber-webapp/helpers"
	"github.com/ortizdavid/golang-fiber-webapp/models"
)

func GetLoggedUser(c *fiber.Ctx) entities.UserData {
    store := config.GetSessionStore()
    session, _ := store.Get(c)
    userName := helpers.ConvertToString(session.Get("username"))
    password := helpers.ConvertToString(session.Get("password"))
	loggedUser, _ := models.UserModel{}.GetByUserNameAndPassword(userName, password)
    return loggedUser
}

func IsUserNomal(c *fiber.Ctx) bool {
	loggedUser := GetLoggedUser(c)
	return loggedUser.RoleCode == "normal"
}

func IsUserAdmin(c *fiber.Ctx) bool {
	loggedUser := GetLoggedUser(c)
	return loggedUser.RoleCode == "admin"
}

func PaginationHandler(c *fiber.Ctx) error {
	var pagination helpers.Pagination
	pageNumber := pagination.GetPageNumber(c, "page")
	return c.Redirect(fmt.Sprintf("/?page=%d", pageNumber))
}

