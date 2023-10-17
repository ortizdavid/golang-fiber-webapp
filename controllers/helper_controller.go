package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-fiber-webapp/config"
	"github.com/ortizdavid/golang-fiber-webapp/entities"
	"github.com/ortizdavid/golang-fiber-webapp/helpers"
	"github.com/ortizdavid/golang-fiber-webapp/models"
)

func GetLoggedUser(ctx *fiber.Ctx) entities.UserData {
    store := config.GetSessionStore()
    session, _ := store.Get(ctx)
    userName := helpers.ConvertToString(session.Get("username"))
    password := helpers.ConvertToString(session.Get("password"))
    return models.UserModel{}.GetByUserNameAndPassword(userName, password)
}

func IsUserNomal(ctx *fiber.Ctx) bool {
	loggedUser := GetLoggedUser(ctx)
	return loggedUser.RoleCode == "normal"
}

func IsUserAdmin(ctx *fiber.Ctx) bool {
	loggedUser := GetLoggedUser(ctx)
	return loggedUser.RoleCode == "admin"
}

func PaginationHandler(ctx *fiber.Ctx) error {
	var pagination helpers.Pagination
	pageNumber := pagination.GetPageNumber(ctx, "page")
	return ctx.Redirect(fmt.Sprintf("/?page=%d", pageNumber))
}
