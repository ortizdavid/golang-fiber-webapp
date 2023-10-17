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

type TaskController struct {
}

var loggerTask = config.NewLogger("task-logs.log")

func (task TaskController) RegisterRoutes(router *fiber.App) {
	group := router.Group("/tasks")
	group.Get("/", task.index)
	group.Get("/page/:pageNumber", PaginationHandler)
	group.Get("/add", task.addForm)
	group.Post("/add", task.add)
	group.Get("/:id/details", task.details)
	group.Get("/:id/edit", task.editForm)
	group.Post("/:id/edit", task.edit)
	group.Get("/search", task.searchForm)
	group.Post("/search", task.search)
	group.Get("/:id/add-attachment", task.addAttachmentForm)
	group.Post("/:id/add-attachment", task.addAttachment)
	group.Get("/:id/view-attachment", task.viewAttachment)
}

func (TaskController) index(ctx *fiber.Ctx) error {
	
	var pagination helpers.Pagination
	var taskModel models.TaskModel
	var tasks []entities.TaskData
	var count int

	loggedUser := GetLoggedUser(ctx)
	pageNumber := pagination.GetPageNumber(ctx, "page")
	itemsPerPage := config.ItemsPerPage()
	startIndex := pagination.CalculateStartIndex(pageNumber, itemsPerPage)
	tasks = taskModel.FindAllDataLimit(startIndex, itemsPerPage)
	count = int(taskModel.Count())
	totalPages := pagination.CalculateTotalPages(count, itemsPerPage)

	if pageNumber > totalPages {
		return ctx.Render("error/pagination", fiber.Map{
			"Title": "Tasks",
			"TotalPages": totalPages, 
			"LoggedUser": loggedUser,
		})
	}
	if loggedUser.RoleCode != "normal" {
		return ctx.Render("task/index", fiber.Map{
			"Title": "Tasks",
			"Tasks": tasks,
			"Pagination": helpers.NewPaginationRender(pageNumber),
			"Count": count,
			"LoggedUser": loggedUser,
		})
	}
	userTasks := taskModel.FindAllDataByUserIdLimit(loggedUser.UserId, startIndex, itemsPerPage)
	count = int(taskModel.CountByUser(loggedUser.UserId))
	return ctx.Render("task/my-tasks", fiber.Map{
		"Title": "My Tasks",
		"Tasks": userTasks,
		"Pagination": helpers.NewPaginationRender(pageNumber),
		"Count": count,
		"LoggedUser": loggedUser,
	})
	
}

func (TaskController) details(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	task := models.TaskModel{}.GetDataByUniqueId(id)
	return ctx.Render("task/details", fiber.Map{
		"Title": "Task Details",
		"Task": task,
		"LoggedUser": GetLoggedUser(ctx),
	})
}

func (TaskController) addForm(ctx *fiber.Ctx) error {
	return ctx.Render("task/add", fiber.Map{
		"Title": "Add Tasks",
		"Complexities": models.TaskComplexityModel{}.FindAll(),
		"Statuses": models.TaskStatusModel{}.FindAll(),
		"LoggedUser": GetLoggedUser(ctx),
	})
}

func (TaskController) add(ctx *fiber.Ctx) error {
	loggedUser := GetLoggedUser(ctx)
	taskName := ctx.FormValue("task_name")
	statusId := ctx.FormValue("status_id")
	complexityId := ctx.FormValue("complexity_id")
	description := ctx.FormValue("description")
	startDate := ctx.FormValue("start_date")
	endDate := ctx.FormValue("end_date")

	var taskModel models.TaskModel
	task := entities.Task{
		TaskId:       0,
		UserId:       loggedUser.UserId,
		StatusId:     helpers.ConvertToInt(statusId),
		ComplexityId: helpers.ConvertToInt(complexityId),
		TaskName:     taskName,
		StartDate:    helpers.StringToDate(startDate),
		EndDate:      helpers.StringToDate(endDate),
		Description:  description,
		Attachment:   "",
		UniqueId:     helpers.GenerateUUID(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	taskModel.Create(task)
	loggerTask.Info(fmt.Sprintf("User '%s' added Task '%s'", loggedUser.UserName, taskName))
	return ctx.Redirect("/tasks")
}


func (TaskController) editForm(ctx *fiber.Ctx) error {
	loggedUser := GetLoggedUser(ctx)
	id := ctx.Params("id")
	task := models.TaskModel{}.GetDataByUniqueId(id)
	return ctx.Render("task/edit", fiber.Map{
		"Title": "Edit Task",
		"Task": task,
		"Statuses": models.TaskStatusModel{}.FindAll(),
		"Complexities": models.TaskComplexityModel{}.FindAll(),
		"LoggedUser": loggedUser,
	})
}

func (TaskController) edit(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	taskName := ctx.FormValue("task_name")
	statusId := ctx.FormValue("status_id")
	complexityId := ctx.FormValue("complexity_id")
	description := ctx.FormValue("description")
	startDate := ctx.FormValue("start_date")
	endDate := ctx.FormValue("end_date")

	var taskModel models.TaskModel
	task := taskModel.FindByUniqueId(id)
	task.TaskName = taskName
	task.StatusId = helpers.ConvertToInt(statusId)
	task.ComplexityId = helpers.ConvertToInt(complexityId)
	task.Description = description
	task.StartDate = helpers.StringToDate(startDate)
	task.EndDate = helpers.StringToDate(endDate)
	task.UpdatedAt = time.Now()
	taskModel.Update(task)
	loggerTask.Info(fmt.Sprintf("Task '%s' Added ", taskName))
	return ctx.Redirect(fmt.Sprintf("/tasks/%s/details", id))
}

func (TaskController) searchForm(ctx *fiber.Ctx) error {
	return ctx.Render("task/search", fiber.Map{
		"Title": "Search Tasks",
		"LoggedUser": GetLoggedUser(ctx),
	})
}

func (TaskController) search(ctx *fiber.Ctx) error {
	param := ctx.FormValue("search_param")
	results := models.TaskModel{}.Search(param)
	count := len(results)
	loggedUser := GetLoggedUser(ctx)
	loggerTask.Info(fmt.Sprintf("User '%s' Searched for Task '%v' and Found %d results", loggedUser.UserName, param, count))
	return ctx.Render("task/search-results", fiber.Map{
		"Title": "Results",
		"Results": results,
		"Param": param,
		"Count": count,
		"LoggedUser": loggedUser,
	})
}

func (TaskController) addAttachmentForm(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	task := models.TaskModel{}.FindByUniqueId(id)
	return ctx.Render("task/add-attachment", fiber.Map{
		"Title": "Add Attachment",
		"Task": task,
		"LoggedUser": GetLoggedUser(ctx),
	})
}

func (TaskController) addAttachment(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	attachment, _ := helpers.UploadFile(ctx, "attachment", "document", config.UploadDocumentPath())
	loggedUser := GetLoggedUser(ctx)

	var taskModel models.TaskModel
	task := taskModel.FindByUniqueId(id)
	task.Attachment = attachment
	task.UpdatedAt = time.Now()
	taskModel.Update(task)
	loggerTask.Info(fmt.Sprintf("User '%s' added attachment for task '%s' ", loggedUser.UserName, task.TaskName))
	return ctx.Redirect("/tasks/"+id+"/details")
}


func (TaskController) viewAttachment(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	task := models.TaskModel{}.FindByUniqueId(id)
	return ctx.SendFile("./static/uploads/docs/"+task.Attachment)
}