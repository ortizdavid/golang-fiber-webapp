package models

import (
	"github.com/ortizdavid/golang-fiber-webapp/config"
	"github.com/ortizdavid/golang-fiber-webapp/entities"
	"gorm.io/gorm"
)

type RoleModel struct {
}


func (RoleModel) Create(role entities.Role) *gorm.DB {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	return db.Create(&role)
}

func (RoleModel) FindAll() []entities.Role {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	roles := []entities.Role{}
	db.Find(&roles)
	return roles
}

func (RoleModel) Update(role entities.Role) *gorm.DB {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	return db.Save(&role)
}

func (RoleModel) FindById(id int) entities.Role {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var role entities.Role
	db.First(&role, id)
	return role
}

func (RoleModel) FindByName(name string) entities.Role {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var role entities.Role
	db.Where("role_name=?", name).First(&role)
	return role
}

func (RoleModel) Count() int64 {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var count int64
	db.Table("roles").Count(&count)
	return count
}
