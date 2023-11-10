package controllers

import (
	"fmt"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-fiber-webapp/config"
	"github.com/ortizdavid/golang-fiber-webapp/entities"
	"github.com/ortizdavid/golang-fiber-webapp/helpers"
	"github.com/ortizdavid/golang-fiber-webapp/models"
)

type UserController struct {
}

type UserValidation struct {
	Username string `form:"username" validate:"required,min=3"`
	Password    string `form:"password" validate:"required"`
	RoleId    string `form:"role_id" validate:"required"`
}

var loggerUser = config.NewLogger("user-logs.log")

func (user UserController) RegisterRoutes(router *fiber.App) {
	group := router.Group("/users")
	group.Get("/", user.index)
	group.Get("/page/:pageNumber", PaginationHandler)
	group.Get("/add", user.addForm)
	group.Post("/add", user.add)
	group.Get("/:id/details", user.details)
	group.Get("/:id/edit", user.editForm)
	group.Post("/:id/edit", user.edit)
	group.Get("/:id/deactivate", user.deactivateForm)
	group.Post("/:id/deactivate", user.deactivate)
	group.Get("/:id/activate", user.activateForm)
	group.Post("/:id/activate", user.activate)
	group.Get("/add-image", user.addImageForm)
	group.Post("/add-image", user.addImage)
	group.Get("/search", user.searchForm)
	group.Post("/search", user.search)
}

func (UserController) index(ctx *fiber.Ctx) error {
	var pagination helpers.Pagination
	var userModel models.UserModel
	
	loggedUser := GetLoggedUser(ctx)
	pageNumber := pagination.GetPageNumber(ctx, "page")
	itemsPerPage := config.ItemsPerPage()
	startIndex := pagination.CalculateStartIndex(pageNumber, itemsPerPage)
	users, _ := userModel.FindAllDataLimit(startIndex, itemsPerPage)
	countUsers, _ := userModel.Count()
	count := int(countUsers)
	totalPages := pagination.CalculateTotalPages(count, itemsPerPage)

	if pageNumber > totalPages {
		return ctx.Render("error/pagination", fiber.Map{
			"Title": "Tasks",
			"TotalPages": totalPages, 
			"LoggedUser": loggedUser,
		})
	}
	return ctx.Render("user/index", fiber.Map{
		"Title": "Users",
		"Users": users,
		"Pagination": helpers.NewPaginationRender(pageNumber),
		"Count": count,
		"PageNumber": pageNumber,
		"TotalPages": totalPages,
		"LoggedUser": loggedUser,
	})
}

func (UserController) details(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, _ := models.UserModel{}.GetDataByUniqueId(id)
	return ctx.Render("user/details", fiber.Map{
		"Title": "User Details",
		"User": user,
		"LoggedUser": GetLoggedUser(ctx),
	})
}

func (UserController) addForm(ctx *fiber.Ctx) error {
	roles, _ := models.RoleModel{}.FindAll()
	return ctx.Render("user/add", fiber.Map{
		"Title": "Add User",
		"Roles": roles,
		"LoggedUser": GetLoggedUser(ctx),
	})
}

func (UserController) add(ctx *fiber.Ctx) error {
	userName := ctx.FormValue("username")
	password := ctx.FormValue("password")
	roleId := ctx.FormValue("role_id")

	var userModel models.UserModel
	user := entities.User{
		UserId:    0,
		RoleId:    helpers.ConvertToInt(roleId),
		UserName:  userName,
		Password:  helpers.HashPassword(password),
		Active:    "Yes",
		Image:     "",
		UniqueId:  helpers.GenerateUUID(),
		Token:     helpers.GenerateRandomToken(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	userModel.Create(user)
	loggerUser.Info(fmt.Sprintf("User '%s' added successfully", userName))
	return ctx.Redirect("/users")
}

func (UserController) editForm(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, _ := models.UserModel{}.GetDataByUniqueId(id)
	roles, _ := models.RoleModel{}.FindAll()
	return ctx.Render("user/edit", fiber.Map{
		"Title": "Edit User",
		"Roles": roles,
		"User": user,
		"LoggedUser": GetLoggedUser(ctx),
	})
}

func (UserController) edit(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	roleId := ctx.FormValue("role_id")
	var userModel models.UserModel
	user, _ := userModel.FindByUniqueId(id)
	user.RoleId = helpers.ConvertToInt(roleId)
	user.UpdatedAt = time.Now()
	user.Token = helpers.GenerateRandomToken()
	userModel.Update(user)
	loggerUser.Info(fmt.Sprintf("User '%s' Updated successfully", user.UserName))
	return ctx.Redirect("/users")
}


func (UserController) searchForm(ctx *fiber.Ctx) error {
	return ctx.Render("user/search", fiber.Map{
		"Title": "Search Users",
		"LoggedUser": GetLoggedUser(ctx),
	})
}

func (UserController) search(ctx *fiber.Ctx) error {
	param := ctx.FormValue("search_param")
	results, _ := models.UserModel{}.Search(param)
	count := len(results)
	loggedUser := GetLoggedUser(ctx)
	loggerUser.Info(fmt.Sprintf("User '%s' Searched for User '%v' and Found %d results", loggedUser.UserName, param, count))
	return ctx.Render("user/search-results", fiber.Map{
		"Title": "Results",
		"Results": results,
		"Param": param,
		"Count": count,
		"LoggedUser": loggedUser,
	})
}

func (UserController) deactivateForm(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, _ := models.UserModel{}.GetDataByUniqueId(id)
	return ctx.Render("user/deactivate", fiber.Map{
		"Title": "Deactivate User",
		"User": user,
		"LoggedUser": GetLoggedUser(ctx),
	})
}

func (UserController) deactivate(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var userModel models.UserModel
	user, _ := userModel.FindByUniqueId(id)
	user.Active = "No"
	user.UpdatedAt = time.Now()
	userModel.Update(user)
	loggerUser.Info(fmt.Sprintf("User '%s' deactivated successfuly", user.UserName))
	return ctx.Redirect("/users/"+id+"/details")
}

func (UserController) activateForm(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, _ := models.UserModel{}.GetDataByUniqueId(id)
	return ctx.Render("user/activate", fiber.Map{
		"Title": "Activate User",
		"User": user,
		"LoggedUser": GetLoggedUser(ctx),
	})
}

func (UserController) activate(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var userModel models.UserModel
	user, _ := userModel.FindByUniqueId(id)
	user.Active = "Yes"
	user.UpdatedAt = time.Now()
	user.Token = helpers.GenerateRandomToken()
	userModel.Update(user)
	loggerUser.Info(fmt.Sprintf("User '%s' activated sucessfully", user.UserName))
	return ctx.Redirect("/users/"+id+"/details")
}

func (UserController) addImageForm(ctx *fiber.Ctx) error {
	return ctx.Render("user/add-image", fiber.Map{
		"Title": "Add Image",
		"LoggedUser": GetLoggedUser(ctx),
	})
}

func (UserController) addImage(ctx *fiber.Ctx) error {
	userImage, _ := helpers.UploadFile(ctx, "user_image", "image", config.UploadImagePath())
	loggedUser := GetLoggedUser(ctx)
	var userModel models.UserModel
	user, _ := userModel.FindById(loggedUser.UserId)
	user.Image = userImage
	user.UpdatedAt = time.Now()
	userModel.Update(user)
	loggerUser.Info(fmt.Sprintf("User '%s' added image", user.UserName))
	return ctx.Redirect("/user-data")
}
