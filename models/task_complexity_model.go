package models

import (
	"gorm.io/gorm"
	"github.com/ortizdavid/golang-fiber-webapp/config"
	"github.com/ortizdavid/golang-fiber-webapp/entities"
)

type TaskComplexityModel struct {
}

func (TaskComplexityModel) Create(complexity entities.TaskComplexity) *gorm.DB {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	return db.Create(&complexity)
}

func (TaskComplexityModel) FindAll() []entities.TaskComplexity {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	complexities := []entities.TaskComplexity{}
	db.Find(&complexities)
	return complexities
}

func (TaskComplexityModel) Update(complexity entities.TaskComplexity) *gorm.DB {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	return db.Save(&complexity)
}

func (TaskComplexityModel) FindById(id int) entities.TaskComplexity {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var complexity entities.TaskComplexity
	db.First(&complexity, id)
	return complexity
}

func (TaskComplexityModel) FindByCode(code string) entities.TaskComplexity {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var complexity entities.TaskComplexity
	db.Where("code=?", code).First(&complexity)
	return complexity
}

func (TaskComplexityModel) Count() int64 {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var count int64
	db.Table("task_complexity").Count(&count)
	return count
}
