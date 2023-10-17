package models

import (
	"github.com/ortizdavid/golang-fiber-webapp/config"
	"github.com/ortizdavid/golang-fiber-webapp/entities"
	"gorm.io/gorm"
)

type UserModel struct {
	LastInsertId int
}

func (userModel *UserModel) Create(user entities.User) *gorm.DB {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Create(&user)
	userModel.LastInsertId = user.UserId
	return result
}

func (UserModel) FindAll() []entities.User {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	users := []entities.User{}
	db.Find(&users)
	return users
}

func (UserModel) Update(user entities.User) *gorm.DB {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	return db.Save(&user)
}

func (UserModel) FindById(id int) entities.User {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var user entities.User
	db.First(&user, id)
	return user
}

func (UserModel) FindByUniqueId(uniqueId string) entities.User {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var user entities.User
	db.First(&user, "unique_id=?", uniqueId)
	return user
}

func (UserModel) Search(param interface{}) []entities.UserData {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var users []entities.UserData
	db.Raw("SELECT * FROM view_user_data WHERE user_name=? OR role_name=?", param, param).Scan(&users)
	return users
}

func (UserModel) Count() int64 {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var count int64
	db.Table("users").Count(&count)
	return count
}

func (UserModel) FindAllOrdered() []entities.User {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	users := []entities.User{}
	db.Order("user_name ASC").Find(&users)
	return users
}

func (UserModel) GetDataById(id int) entities.UserData {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var userData entities.UserData
	db.Raw("SELECT * FROM view_user_data WHERE user_id=?", id).Scan(&userData)
	return userData
}

func (UserModel) GetDataByUniqueId(uniqueId string) entities.UserData {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var userData entities.UserData
	db.Raw("SELECT * FROM view_user_data WHERE unique_id=?", uniqueId).Scan(&userData)
	return userData
}

func (UserModel) FindAllData() []entities.UserData {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var users []entities.UserData
	db.Raw("SELECT * FROM view_user_data").Scan(&users)
	return users
}

func (UserModel) FindAllDataLimit(start int, end int) []entities.UserData {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var users []entities.UserData
	db.Raw("SELECT * FROM view_user_data LIMIT ?, ?", start, end).Scan(&users)
	return users
}

func (UserModel) Exists(userName string, password string) bool {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var user entities.User
	db.Where("user_name=? AND password=?", userName, password).Find(&user)
	return user.UserId != 0
}

func (UserModel) ExistsActive(userName string, password string) bool {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var user entities.User
	db.Where("user_name=? AND password=? AND active='Yes'", userName, password).Find(&user)
	return user.UserId != 0
}

func (UserModel) ExistsActiveUser(userName string) bool {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var user entities.User
	db.Where("user_name=? AND active='Yes'", userName).Find(&user)
	return user.UserId != 0
}

func (UserModel) GetByUserNameAndPassword(userName string, password string) entities.UserData {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var userData entities.UserData
	db.Raw("SELECT * FROM view_user_data WHERE user_name=? AND password=?", userName, password).Scan(&userData)
	return userData
}

func (UserModel) GetByUserName(userName string) entities.UserData {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var userData entities.UserData
	db.Raw("SELECT * FROM view_user_data WHERE user_name=?", userName).Scan(&userData)
	return userData
}

func (UserModel) FindAllByStatus(statusName string) []entities.UserData {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	users := []entities.UserData{}
	db.Raw("SELECT * FROM view_user_data WHERE active=?",  statusName).Find(&users)
	return users
}

func (UserModel) FindAllByRole(roleName string) []entities.UserData {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	users := []entities.UserData{}
	db.Raw("SELECT * FROM view_user_data WHERE role_name=?",  roleName).Find(&users)
	return users
}

func (UserModel) FindInactiveByRole(roleName string) []entities.UserData {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	users := []entities.UserData{}
	db.Raw("SELECT * FROM view_user_data WHERE role_name=? and active='No'",  roleName).Find(&users)
	return users
}
