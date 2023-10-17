package entities

type TaskComplexity struct {
	ComplexityId	int `gorm:"primaryKey;autoIncrement"`
	ComplexityName	string `gorm:"column:complexity_name;type:varchar(100)"`
	Code			string `gorm:"column:code;type:varchar(20)"`
}

func (TaskComplexity) TableName() string {
	return "task_complexity"
}