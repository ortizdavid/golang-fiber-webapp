package entities

type TaskStatus struct {
	StatusId	int `gorm:"primaryKey;autoIncrement"`
	StatusName	string `gorm:"column:status_name;type:varchar(100)"`
	Code		string `gorm:"column:code;type:varchar(20)"`
}

func (TaskStatus) TableName() string {
	return "task_status"
}