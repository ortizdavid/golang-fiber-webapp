package main

import (
	"fmt"	
	"time"
	"log/slog"
	"github.com/ortizdavid/golang-fiber-webapp/entities"
	"github.com/ortizdavid/golang-fiber-webapp/helpers"
	"github.com/ortizdavid/golang-fiber-webapp/models"
)

func main() {
	var userName, password, roleName string
	var roleId int
	var roleModel models.RoleModel
	var userModel models.UserModel

	fmt.Println("Create Users")
	fmt.Print("Role: \n\t[1]-Admin\n\t[2]-Normal: ")
	fmt.Scanln(&roleId)
	fmt.Print("Username: ")
	fmt.Scanln(&userName)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	if roleId < 1 || roleId > 2 {
		roleId = 1
	}
	roleName = roleModel.FindById(roleId).RoleName

	user := entities.User{
		UserId:    0,
		RoleId:    roleId,
		UserName:  userName,
		Password:  helpers.HashPassword(password),
		Active:    "Yes",
		Image:     "",
		UniqueId:  helpers.GenerateUUID(),
		CreatedAt: time.Now(),
		UpdatedAt:  time.Now(),
	}
	userModel.Create(user)
	slog.Info("User '%s' Created with role '%s'", userName, roleName)
}