package models

type StatisticCount struct {
	Users 				int64
	Tasks 				int64
	PendingTasks		int64
	CompletedTasks 		int64
	InProgressTasks 	int64
	BlockedTasks 		int64
	CanceledTasks		int64
}

func GetStatisticsCount() StatisticCount {
	return StatisticCount{
		Users:           UserModel{}.Count(),
		Tasks:           TaskModel{}.Count(),
		PendingTasks:    TaskModel{}.CountByStatus("pending"),
		CompletedTasks:  TaskModel{}.CountByStatus("completed"),
		InProgressTasks: TaskModel{}.CountByStatus("in-progress"),
		BlockedTasks:    TaskModel{}.CountByStatus("blocked"),
		CanceledTasks:   TaskModel{}.CountByStatus("canceled"),
	}
}

func GetStatisticsCountByUser(userId int) StatisticCount {
	return StatisticCount{
		Users:           0,
		Tasks:           TaskModel{}.CountByUser(userId),
		PendingTasks:    TaskModel{}.CountByStatusAndUser("pending", userId),
		CompletedTasks:  TaskModel{}.CountByStatusAndUser("completed", userId),
		InProgressTasks: TaskModel{}.CountByStatusAndUser("in-progress", userId),
		BlockedTasks:    TaskModel{}.CountByStatusAndUser("blocked", userId),
		CanceledTasks:   TaskModel{}.CountByStatusAndUser("canceled", userId),
	}
}

