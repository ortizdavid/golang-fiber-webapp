package models

import (
	"gorm.io/gorm"
	"github.com/ortizdavid/golang-fiber-webapp/config"
	"github.com/ortizdavid/golang-fiber-webapp/entities"
)

type TaskStatusModel struct {
}

func (TaskStatusModel) Create(status entities.TaskStatus) (*gorm.DB, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Create(&status)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (TaskStatusModel) FindAll() ([]entities.TaskStatus, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	statuses := []entities.TaskStatus{}
	result := db.Find(&statuses)
	if result.Error != nil {
		return nil, result.Error
	}
	return statuses, nil
}

func (TaskStatusModel) Update(status entities.TaskStatus) (*gorm.DB, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Save(&status)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (TaskStatusModel) FindById(id int) (entities.TaskStatus, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var status entities.TaskStatus
	result := db.First(&status, id)
	if result.Error != nil {
		return entities.TaskStatus{}, result.Error
	}
	return status, nil
}

func (TaskStatusModel) FindByCode(code string) (entities.TaskStatus, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var status entities.TaskStatus
	result := db.Where("code=?", code).First(&status)
	if result.Error != nil {
		return entities.TaskStatus{}, result.Error
	}
	return status, nil
}

func (TaskStatusModel) Count() (int64, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var count int64
	result := db.Table("task_status").Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}
