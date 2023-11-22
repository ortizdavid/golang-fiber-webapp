package models

import (
	"io"
	"fmt"
	"time"
	"encoding/csv"
	"github.com/ortizdavid/golang-fiber-webapp/entities"
	"github.com/ortizdavid/golang-fiber-webapp/helpers"
)

const (
	statusIndex			= 0
	complexityIndex		= 1
	taskNameIndex		= 2
	startDateIndex		= 3
	endDateIndex		= 4
	descriptionIndex	= 5
)

func ParseTaskFromCSV(reader *csv.Reader, userId int) ([]entities.Task, error) {
	var tasks []entities.Task
	var lineNumber int

	for {
		record, err := reader.Read()
		lineNumber++
		
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("error reading CSV record %v on line %d", err, lineNumber)
		}

		if len(record) != 6 {
			return nil, fmt.Errorf("invalid CSV record %v", err)
		}

		task := entities.Task {
			TaskId:       0,
			UserId:       userId,
			StatusId:     helpers.ConvertToInt(record[statusIndex]),
			ComplexityId: helpers.ConvertToInt(record[complexityIndex]),
			TaskName:     record[taskNameIndex],
			StartDate:    record[startDateIndex],
			EndDate:      record[endDateIndex],
			Description:  record[descriptionIndex],
			Attachment:   "",
			UniqueId:     helpers.GenerateUUID(),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}


func SkipCSVHeader(reader *csv.Reader) error {
	_, err := reader.Read()
	return err
}