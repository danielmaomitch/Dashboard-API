package TaskList

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

func unmarshalTasks(query *dynamodb.QueryOutput) []Task {

	taskList := []Task{}
	for _, dbTask := range query.Items {
		task := Task{}
		if err := dynamodbattribute.UnmarshalMap(dbTask, &task); err != nil {
			log.Fatalf("Error unmarshalling tasks: %s", err)
		}

		taskList = append(taskList, task)
	}
	return taskList
}

func AddRecord(t Task) string {
	t.SK = "Task_" + uuid.NewString()
	t.CreatedAt = time.Now().Format(time.DateOnly)
	DBCreate(t)

	return t.SK
}

func GetRecord() []Task {
	result := DBRead("twiggs", "Task")
	taskList := unmarshalTasks(result)

	return taskList
}

func DelRecord(userID string, id string) (string, error) {
	SK := "Task_" + id
	if err := DBDelete(userID, SK); err != nil {
		return "", err
	}

	return SK, nil
}
