package worker

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type TaskResult struct {
	TaskID uint   `json:"task_id"`
	Status string `json:"status"`
}

func SendTaskResult(taskID uint, status string) {
	result := TaskResult{
		TaskID: taskID,
		Status: status,
	}
	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Println("Failed to marshal task result:", err)
		return
	}

	_, err = http.Post("http://scheduler:8080/tasks/result", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Failed to send task result:", err)
	}
}
