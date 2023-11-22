package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"time"

	"github.com/ortizdavid/golang-fiber-webapp/entities"
	"github.com/ortizdavid/golang-fiber-webapp/helpers"
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
			StatusId:     helpers.ConvertToInt(record[0]),
			ComplexityId: helpers.ConvertToInt(record[1]),
			TaskName:     record[2],
			StartDate:    record[3],
			EndDate:      record[4],
			Description:  record[5],
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