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

func (UserController) index(c *fiber.Ctx) error {
	var pagination helpers.Pagination
	var userModel models.UserModel
	
	loggedUser := GetLoggedUser(c)
	pageNumber := pagination.GetPageNumber(c, "page")
	itemsPerPage := config.ItemsPerPage()
	startIndex := pagination.CalculateStartIndex(pageNumber, itemsPerPage)
	users, _ := userModel.FindAllDataLimit(startIndex, itemsPerPage)
	countUsers, _ := userModel.Count()
	count := int(countUsers)
	totalPages := pagination.CalculateTotalPages(count, itemsPerPage)

	if count > 0 && pageNumber > totalPages {
		return c.Status(500).Render("error/pagination", fiber.Map{
			"Title": "Tasks",
			"TotalPages": totalPages, 
			"LoggedUser": loggedUser,
		})
	}
	return c.Render("user/index", fiber.Map{
		"Title": "Users",
		"Users": users,
		"Pagination": helpers.NewPaginationRender(pageNumber),
		"Count": count,
		"PageNumber": pageNumber,
		"TotalPages": totalPages,
		"LoggedUser": loggedUser,
	})
}

func (UserController) details(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := models.UserModel{}.GetDataByUniqueId(id)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.Render("user/details", fiber.Map{
		"Title": "User Details",
		"User": user,
		"LoggedUser": GetLoggedUser(c),
	})
}

func (UserController) addForm(c *fiber.Ctx) error {
	roles, err := models.RoleModel{}.FindAll()
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.Render("user/add", fiber.Map{
		"Title": "Add User",
		"Roles": roles,
		"LoggedUser": GetLoggedUser(c),
	})
}

func (UserController) add(c *fiber.Ctx) error {
	userName := c.FormValue("username")
	password := c.FormValue("password")
	roleId := c.FormValue("role_id")

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
	_, err := userModel.Create(user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	loggerUser.Info(fmt.Sprintf("User '%s' added successfully", userName))
	return c.Redirect("/users")
}

func (UserController) editForm(c *fiber.Ctx) error {
	id := c.Params("id")
	user, _ := models.UserModel{}.GetDataByUniqueId(id)
	roles, _ := models.RoleModel{}.FindAll()
	return c.Render("user/edit", fiber.Map{
		"Title": "Edit User",
		"Roles": roles,
		"User": user,
		"LoggedUser": GetLoggedUser(c),
	})
}

func (UserController) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	roleId := c.FormValue("role_id")
	var userModel models.UserModel
	user, _ := userModel.FindByUniqueId(id)
	user.RoleId = helpers.ConvertToInt(roleId)
	user.UpdatedAt = time.Now()
	user.Token = helpers.GenerateRandomToken()
	_, err := userModel.Update(user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	loggerUser.Info(fmt.Sprintf("User '%s' Updated successfully", user.UserName))
	return c.Redirect("/users")
}


func (UserController) searchForm(c *fiber.Ctx) error {
	return c.Render("user/search", fiber.Map{
		"Title": "Search Users",
		"LoggedUser": GetLoggedUser(c),
	})
}

func (UserController) search(c *fiber.Ctx) error {
	param := c.FormValue("search_param")
	results, err := models.UserModel{}.Search(param)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	count := len(results)
	loggedUser := GetLoggedUser(c)
	loggerUser.Info(fmt.Sprintf("User '%s' Searched for User '%v' and Found %d results", loggedUser.UserName, param, count))
	return c.Render("user/search-results", fiber.Map{
		"Title": "Results",
		"Results": results,
		"Param": param,
		"Count": count,
		"LoggedUser": loggedUser,
	})
}

func (UserController) deactivateForm(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := models.UserModel{}.GetDataByUniqueId(id)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.Render("user/deactivate", fiber.Map{
		"Title": "Deactivate User",
		"User": user,
		"LoggedUser": GetLoggedUser(c),
	})
}

func (UserController) deactivate(c *fiber.Ctx) error {
	id := c.Params("id")
	var userModel models.UserModel
	user, err := userModel.FindByUniqueId(id)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	user.Active = "No"
	user.UpdatedAt = time.Now()
	userModel.Update(user)
	loggerUser.Info(fmt.Sprintf("User '%s' deactivated successfuly", user.UserName))
	return c.Redirect("/users/"+id+"/details")
}

func (UserController) activateForm(c *fiber.Ctx) error {
	id := c.Params("id")
	user, _ := models.UserModel{}.GetDataByUniqueId(id)
	return c.Render("user/activate", fiber.Map{
		"Title": "Activate User",
		"User": user,
		"LoggedUser": GetLoggedUser(c),
	})
}

func (UserController) activate(c *fiber.Ctx) error {
	id := c.Params("id")
	var userModel models.UserModel
	user, err := userModel.FindByUniqueId(id)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	user.Active = "Yes"
	user.UpdatedAt = time.Now()
	user.Token = helpers.GenerateRandomToken()
	userModel.Update(user)
	loggerUser.Info(fmt.Sprintf("User '%s' activated sucessfully", user.UserName))
	return c.Redirect("/users/"+id+"/details")
}

func (UserController) addImageForm(c *fiber.Ctx) error {
	return c.Render("user/add-image", fiber.Map{
		"Title": "Add Image",
		"LoggedUser": GetLoggedUser(c),
	})
}

func (UserController) addImage(c *fiber.Ctx) error {
	userImage, _ := helpers.UploadFile(c, "user_image", "image", config.UploadImagePath())
	loggedUser := GetLoggedUser(c)
	var userModel models.UserModel
	user, err := userModel.FindById(loggedUser.UserId)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	user.Image = userImage
	user.UpdatedAt = time.Now()
	userModel.Update(user)
	loggerUser.Info(fmt.Sprintf("User '%s' added image", user.UserName))
	return c.Redirect("/user-data")
}
