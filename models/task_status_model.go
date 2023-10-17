package models

import (
	"github.com/ortizdavid/golang-fiber-webapp/config"
	"github.com/ortizdavid/golang-fiber-webapp/entities"
	"gorm.io/gorm"
)

type TaskStatusModel struct {
}

func (TaskStatusModel) Create(status entities.TaskStatus) *gorm.DB {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	return db.Create(&status)
}

func (TaskStatusModel) FindAll() []entities.TaskStatus {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	statuses := []entities.TaskStatus{}
	db.Find(&statuses)
	return statuses
}

func (TaskStatusModel) Update(status entities.TaskStatus) *gorm.DB {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	return db.Save(&status)
}

func (TaskStatusModel) FindById(id int) entities.TaskStatus {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var status entities.TaskStatus
	db.First(&status, id)
	return status
}

func (TaskStatusModel) FindByCode(code string) entities.TaskStatus {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var status entities.TaskStatus
	db.Where("code=?", code).First(&status)
	return status
}

func (TaskStatusModel) Count() int64 {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var count int64
	db.Table("task_status").Count(&count)
	return count
}
