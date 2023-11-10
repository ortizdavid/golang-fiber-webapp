package models

type ReportModel struct {
}

type TableReport struct {
	Title  	string
	Rows	interface{}
	Count   int
}

func (ReportModel) GetAllUsers() TableReport {
	rows, _ := UserModel{}.FindAllData()
	count := len(rows)
	return TableReport {
		Title: "Users",
		Rows: rows,
		Count: count,
	}
}

func (ReportModel) GetAllActiveUsers() TableReport {
	rows, _ := UserModel{}.FindAllByStatus("Yes")
	count := len(rows)
	return TableReport {
		Title: "Active Users",
		Rows: rows,
		Count: count,
	}
}

func (ReportModel) GetAllTasks() TableReport {
	rows, _ := TaskModel{}.FindAllData()
	count := len(rows)
	return TableReport {
		Title: "Tasks",
		Rows: rows,
		Count: count,
	}
}

func (ReportModel) GetAllPendingTasks() TableReport {
	rows, _ := TaskModel{}.FindAllDataByStatus("pending")
	count := len(rows)
	return TableReport {
		Title: "Pending Tasks",
		Rows: rows,
		Count: count,
	}
}

func (ReportModel) GetAllCompletedTasks() TableReport {
	rows, _ := TaskModel{}.FindAllDataByStatus("completed")
	count := len(rows)
	return TableReport {
		Title: "Completed Tasks",
		Rows: rows,
		Count: count,
	}
}

func (ReportModel) GetAllInProgressTasks() TableReport {
	rows, _ := TaskModel{}.FindAllDataByStatus("in-progress")
	count := len(rows)
	return TableReport {
		Title: "In Progress",
		Rows: rows,
		Count: count,
	}
}

func (ReportModel) GetAllBlockedTasks() TableReport {
	rows, _ := TaskModel{}.FindAllDataByStatus("blocked")
	count := len(rows)
	return TableReport {
		Title: "Blocked Tasks",
		Rows: rows,
		Count: count,
	}
}

func (ReportModel) GetAllCancelledTasks() TableReport {
	rows, _ := TaskModel{}.FindAllDataByStatus("canceled")
	count := len(rows)
	return TableReport {
		Title: "Cancelled Tasks",
		Rows: rows,
		Count: count,
	}
}
