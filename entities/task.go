package entities

import (
	"time"
)

type Task struct {
	TaskId      	int `gorm:"primaryKey;autoIncrement"`
	UserId    		int `gorm:"column:user_id;type:int"`
	StatusId    	int `gorm:"column:status_id;type:int"`
	ComplexityId    int `gorm:"column:complexity_id;type:int"`
	TaskName    	string `gorm:"column:task_name;type:varchar(100)"`
	StartDate		string `gorm:"column:start_date;type:date"`
	EndDate			string  `gorm:"column:end_date;type:date"`
	Description 	string `gorm:"column:description;type:varchar(300)"`
	Attachment 		string `gorm:"column:attachment;type:varchar(100)"`
	UniqueId 		string `gorm:"column:unique_id;type:varchar(50)"`
	CreatedAt  		time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt  		time.Time `gorm:"column:updated_at;type:datetime"`
}

func (task Task) TableName() string {
	return "tasks"
}

type TaskData struct {
	TaskId			int
	UniqueId		string
	TaskName 		string
	StartDate		string
	EndDate			string
	Description		string
	Attachment		string 
	CreatedAt		string
	UpdatedAt		string
	UserId			int
	UserName		string 
	StatusId		int
	StatusName 		string
	StatusCode 		string
	ComplexityId	int
	ComplexityName	string
}

