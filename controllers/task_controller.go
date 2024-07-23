package controllers

import (
	"encoding/csv"
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
	group.Get("/upload-csv", task.uploadCSVForm)
	group.Post("/upload-csv", task.uploadCSV)
}


func (TaskController) index(c *fiber.Ctx) error {
	
	var pagination helpers.Pagination
	var taskModel models.TaskModel
	var tasks []entities.TaskData
	var count int

	loggedUser := GetLoggedUser(c)
	pageNumber := pagination.GetPageNumber(c, "page")
	itemsPerPage := config.ItemsPerPage()
	startIndex := pagination.CalculateStartIndex(pageNumber, itemsPerPage)
	tasks, _ = taskModel.FindAllDataLimit(startIndex, itemsPerPage)
	countTasks, _ := taskModel.Count()
	count = int(countTasks)
	totalPages := pagination.CalculateTotalPages(count, itemsPerPage)

	if count > 0 && pageNumber > totalPages {
		return c.Status(500).Render("error/pagination", fiber.Map{
			"Title": "Tasks",
			"TotalPages": totalPages, 
			"LoggedUser": loggedUser,
		})
	}
	if loggedUser.RoleCode != "normal" {
		return c.Render("task/index", fiber.Map{
			"Title": "Tasks",
			"Tasks": tasks,
			"Pagination": helpers.NewPaginationRender(pageNumber),
			"Count": count,
			"LoggedUser": loggedUser,
		})
	}
	userTasks, _ := taskModel.FindAllDataByUserIdLimit(loggedUser.UserId, startIndex, itemsPerPage)
	countByUser, _ := taskModel.CountByUser(loggedUser.UserId)
	count = int(countByUser)
	return c.Render("task/my-tasks", fiber.Map{
		"Title": "My Tasks",
		"Tasks": userTasks,
		"Pagination": helpers.NewPaginationRender(pageNumber),
		"Count": count,
		"LoggedUser": loggedUser,
	})
}


func (TaskController) details(c *fiber.Ctx) error {
	id := c.Params("id")
	task, _ := models.TaskModel{}.GetDataByUniqueId(id)
	return c.Render("task/details", fiber.Map{
		"Title": "Task Details",
		"Task": task,
		"LoggedUser": GetLoggedUser(c),
	})
}


func (TaskController) addForm(c *fiber.Ctx) error {
	complexities, _ := models.TaskComplexityModel{}.FindAll()
	statuses, _ := models.TaskStatusModel{}.FindAll()
	return c.Render("task/add", fiber.Map{
		"Title": "Add Tasks",
		"Complexities": complexities,
		"Statuses": statuses,
		"LoggedUser": GetLoggedUser(c),
	})
}


func (TaskController) add(c *fiber.Ctx) error {
	loggedUser := GetLoggedUser(c)
	taskName := c.FormValue("task_name")
	statusId := c.FormValue("status_id")
	complexityId := c.FormValue("complexity_id")
	description := c.FormValue("description")
	startDate := c.FormValue("start_date")
	endDate := c.FormValue("end_date")

	var taskModel models.TaskModel
	task := entities.Task{
		TaskId:       0,
		UserId:       loggedUser.UserId,
		StatusId:     helpers.ConvertToInt(statusId),
		ComplexityId: helpers.ConvertToInt(complexityId),
		TaskName:     taskName,
		StartDate:    startDate,
		EndDate:      endDate,
		Description:  description,
		Attachment:   "",
		UniqueId:     helpers.GenerateUUID(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	_, err := taskModel.Create(task)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	loggerTask.Info(fmt.Sprintf("User '%s' added Task '%s'", loggedUser.UserName, taskName))
	return c.Redirect("/tasks")
}


func (TaskController) editForm(c *fiber.Ctx) error {
	loggedUser := GetLoggedUser(c)
	id := c.Params("id")
	task, _ := models.TaskModel{}.GetDataByUniqueId(id)
	statuses, _ := models.TaskStatusModel{}.FindAll()
	complexities, _ := models.TaskComplexityModel{}.FindAll()
	return c.Render("task/edit", fiber.Map{
		"Title": "Edit Task",
		"Task": task,
		"Statuses": statuses,
		"Complexities": complexities,
		"LoggedUser": loggedUser,
	})
}


func (TaskController) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	taskName := c.FormValue("task_name")
	statusId := c.FormValue("status_id")
	complexityId := c.FormValue("complexity_id")
	description := c.FormValue("description")
	startDate := c.FormValue("start_date")
	endDate := c.FormValue("end_date")

	var taskModel models.TaskModel
	task, _ := taskModel.FindByUniqueId(id)
	task.TaskName = taskName
	task.StatusId = helpers.ConvertToInt(statusId)
	task.ComplexityId = helpers.ConvertToInt(complexityId)
	task.Description = description
	task.StartDate = startDate
	task.EndDate = endDate
	task.UpdatedAt = time.Now()
	_, err := taskModel.Update(task)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	loggerTask.Info(fmt.Sprintf("Task '%s' Added ", taskName))
	return c.Redirect(fmt.Sprintf("/tasks/%s/details", id))
}


func (TaskController) searchForm(c *fiber.Ctx) error {
	return c.Render("task/search", fiber.Map{
		"Title": "Search Tasks",
		"LoggedUser": GetLoggedUser(c),
	})
}


func (TaskController) search(c *fiber.Ctx) error {
	param := c.FormValue("search_param")
	results, err := models.TaskModel{}.Search(param)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	count := len(results)
	loggedUser := GetLoggedUser(c)
	loggerTask.Info(fmt.Sprintf("User '%s' Searched for Task '%v' and Found %d results", loggedUser.UserName, param, count))
	return c.Render("task/search-results", fiber.Map{
		"Title": "Results",
		"Results": results,
		"Param": param,
		"Count": count,
		"LoggedUser": loggedUser,
	})
}


func (TaskController) addAttachmentForm(c *fiber.Ctx) error {
	id := c.Params("id")
	task, _ := models.TaskModel{}.FindByUniqueId(id)
	return c.Render("task/add-attachment", fiber.Map{
		"Title": "Add Attachment",
		"Task": task,
		"LoggedUser": GetLoggedUser(c),
	})
}


func (TaskController) addAttachment(c *fiber.Ctx) error {
	id := c.Params("id")
	attachment, _ := helpers.UploadFile(c, "attachment", "document", config.UploadDocumentPath())
	loggedUser := GetLoggedUser(c)

	var taskModel models.TaskModel
	task, err := taskModel.FindByUniqueId(id)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	task.Attachment = attachment
	task.UpdatedAt = time.Now()
	taskModel.Update(task)
	loggerTask.Info(fmt.Sprintf("User '%s' added attachment for task '%s' ", loggedUser.UserName, task.TaskName))
	return c.Redirect("/tasks/"+id+"/details")
}


func (TaskController) viewAttachment(c *fiber.Ctx) error {
	id := c.Params("id")
	task, err := models.TaskModel{}.FindByUniqueId(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendFile("./static/uploads/docs/"+task.Attachment)
}


func (TaskController) uploadCSVForm(c *fiber.Ctx) error {
	return c.Render("task/upload-csv", fiber.Map{
		"Title": "Upload Tasks from CSV file",
		"LoggedUser": GetLoggedUser(c),
	})
}


func (TaskController) uploadCSV(c *fiber.Ctx) error {
	var taskModel models.TaskModel
	loggedUser := GetLoggedUser(c)

	file, err := c.FormFile("csv_file")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	src, err := file.Open()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer src.Close()

	reader := csv.NewReader(src)
	if err := models.SkipCSVHeader(reader); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	tasksFromCSV, err := models.ParseTaskFromCSV(reader, loggedUser.UserId) // Parsing CSV
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	_, err = taskModel.CreateBatch(tasksFromCSV) // Inserting Batch
	if  err != nil {
		return c.Status(500).SendString(err.Error())
	}

	loggerTask.Info(fmt.Sprintf("User '%s' Uploaded Task from CSV File '%s'", loggedUser.UserName, file.Filename))
	return c.Redirect("/tasks")
}