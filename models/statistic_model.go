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
	countUsers, _ :=  UserModel{}.Count()
	countTasks, _ := TaskModel{}.Count()
	countPending, _ := TaskModel{}.CountByStatus("pending")
	countCompleted, _ := TaskModel{}.CountByStatus("completed")
	countInProgress, _ := TaskModel{}.CountByStatus("in-progress")
	countBlocked, _ := TaskModel{}.CountByStatus("blocked")
	countCanceled, _ := TaskModel{}.CountByStatus("canceled")
	return StatisticCount{
		Users:           countUsers,
		Tasks:           countTasks,
		PendingTasks:    countPending,
		CompletedTasks:  countCompleted,
		InProgressTasks: countInProgress,
		BlockedTasks:    countBlocked,
		CanceledTasks:   countCanceled,
	}
}

func GetStatisticsCountByUser(userId int) StatisticCount {
	countTasks, _ := TaskModel{}.CountByUser(userId)
	countPending, _ := TaskModel{}.CountByStatusAndUser("pending", userId)
	countCompleted, _ := TaskModel{}.CountByStatusAndUser("completed", userId)
	countInProgress, _ := TaskModel{}.CountByStatusAndUser("in-progress", userId)
	countBlocked, _ := TaskModel{}.CountByStatusAndUser("blocked", userId)
	countCanceled, _ := TaskModel{}.CountByStatusAndUser("canceled", userId)
	return StatisticCount{
		Users:           0,
		Tasks:           countTasks,
		PendingTasks:    countPending,
		CompletedTasks:  countCompleted,
		InProgressTasks: countInProgress,
		BlockedTasks:    countBlocked,
		CanceledTasks:   countCanceled,
	}
}

