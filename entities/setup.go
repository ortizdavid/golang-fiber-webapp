package entities

import "github.com/ortizdavid/golang-fiber-webapp/config"

func SetupMigrations() {
	db, _ := config.ConnectDB()
	db.AutoMigrate(&Role{})
	db.AutoMigrate(&User{})	
	db.AutoMigrate(&TaskComplexity{})
	db.AutoMigrate(&TaskStatus{})
	db.AutoMigrate(&Task{})
}