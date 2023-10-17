package models

import (
	"github.com/ortizdavid/golang-fiber-webapp/config"
	"github.com/ortizdavid/golang-fiber-webapp/entities"
	"gorm.io/gorm"
)

type TaskModel struct {
	LastInsertId 	int
}

func (taskModel *TaskModel) Create(task entities.Task) *gorm.DB {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Create(&task)
	taskModel.LastInsertId = task.TaskId
	return result
}

func (TaskModel) FindAll() []entities.Task {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	Tasks := []entities.Task{}
	db.Find(&Tasks)
	return Tasks
}

func (TaskModel) Update(task entities.Task) *gorm.DB {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	return db.Save(&task)
}

func (TaskModel) FindById(id int) entities.Task {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var task entities.Task
	db.First(&task, id)
	return task
}

func (TaskModel) FindByUniqueId(uniqueId string) entities.Task {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var task entities.Task
	db.First(&task, "unique_id=?", uniqueId)
	return task
}

func (TaskModel) FindUserId(userId int) entities.Task {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var task entities.Task
	db.First(&task, "user_id=?", userId)
	return task
}

func (TaskModel) Search(param interface{}) []entities.TaskData {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var tasks []entities.TaskData
	db.Raw("SELECT * FROM view_task_data WHERE task_name=? OR user_name=? OR status_name=? ", param, param, param).Scan(&tasks)
	return tasks
}

func (TaskModel) GetDataById(id int) entities.TaskData {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var taskData entities.TaskData
	db.Raw("SELECT * FROM view_task_data WHERE task_id=?", id).Scan(&taskData)
	return taskData
}

func (TaskModel) GetDataByUniqueId(uniqueId string) entities.TaskData {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var taskData entities.TaskData
	db.Raw("SELECT * FROM view_task_data WHERE unique_id=?", uniqueId).Scan(&taskData)
	return taskData
}

func (TaskModel) FindAllData() []entities.TaskData {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var tasks []entities.TaskData
	db.Raw("SELECT * FROM view_task_data").Scan(&tasks)
	return tasks
}


func (TaskModel) FindAllDataLimit(start int, end int) []entities.TaskData {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var tasks []entities.TaskData
	db.Raw("SELECT * FROM view_task_data LIMIT ?, ?", start, end).Scan(&tasks)
	return tasks
}

func (TaskModel) FindAllDataByTaskId(taskId int) []entities.TaskData {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var tasks []entities.TaskData
	db.Raw("SELECT * FROM view_task_data WHERE task_id=?", taskId).Scan(&tasks)
	return tasks
}

func (TaskModel) FindAllDataByStatus(status string) []entities.TaskData {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var tasks []entities.TaskData
	db.Raw("SELECT * FROM view_task_data WHERE status_name=?", status).Scan(&tasks)
	return tasks
}

func (TaskModel) FindAllDataByStatusLimit(status string, start int, end int) []entities.TaskData {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var tasks []entities.TaskData
	db.Raw("SELECT * FROM view_task_data WHERE status_name=? LIMIT ?, ?", status, start, end).Scan(&tasks)
	return tasks
}

func (TaskModel) FindAllDataByUserId(userId int) []entities.TaskData {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var tasks []entities.TaskData
	db.Raw("SELECT * FROM view_task_data WHERE user_id=?", userId).Scan(&tasks)
	return tasks
}

func (TaskModel) FindAllDataByUserIdLimit(userId int, start int, end int) []entities.TaskData {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var tasks []entities.TaskData
	db.Raw("SELECT * FROM view_task_data WHERE user_id=? LIMIT ?, ?", userId, start, end).Scan(&tasks)
	return tasks
}

func (TaskModel) Count() int64 {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var count int64
	db.Table("tasks").Count(&count)
	return count
}

func (TaskModel) CountByStatus(status string) int64 {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var count int64
	db.Table("view_task_data").Where("status_code=?", status).Count(&count)
	return count
}

func (TaskModel) CountByUser(userId int) int64 {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var count int64
	db.Table("tasks").Where("user_id=?", userId).Count(&count)
	return count
}

func (TaskModel) CountByStatusAndUser(status string, userId int) int64 {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var count int64
	db.Table("view_task_data").Where("status_code=? AND user_id=?", status, userId).Count(&count)
	return count
}
