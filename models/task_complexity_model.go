package models

import (
	"gorm.io/gorm"
	"github.com/ortizdavid/golang-fiber-webapp/config"
	"github.com/ortizdavid/golang-fiber-webapp/entities"
)

type TaskComplexityModel struct {
}

func (TaskComplexityModel) Create(complexity entities.TaskComplexity) (*gorm.DB, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Create(&complexity)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (TaskComplexityModel) FindAll() ([]entities.TaskComplexity, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var complexities []entities.TaskComplexity
	result := db.Find(&complexities)
	if result.Error != nil {
		return nil, result.Error
	}
	return complexities, nil
}

func (TaskComplexityModel) Update(complexity entities.TaskComplexity) (*gorm.DB, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Save(&complexity)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (TaskComplexityModel) FindById(id int) (entities.TaskComplexity, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var complexity entities.TaskComplexity
	result := db.First(&complexity, id)
	if result.Error != nil {
		return entities.TaskComplexity{}, result.Error
	}
	return complexity, nil
}

func (TaskComplexityModel) FindByCode(code string) (entities.TaskComplexity, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var complexity entities.TaskComplexity
	result := db.Where("code=?", code).First(&complexity)
	if result.Error != nil {
		return entities.TaskComplexity{}, result.Error
	}
	return complexity, nil
}

func (TaskComplexityModel) Count() (int64, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var count int64
	result := db.Table("task_complexity").Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}
