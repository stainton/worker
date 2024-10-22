package worker

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/stainton/worker/internal/taskexecutor"
)

type TaskRequest struct {
	TaskID  uint   `json:"task_id"`
	Command string `json:"command"`
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	var taskReq TaskRequest
	err := json.NewDecoder(r.Body).Decode(&taskReq)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	result := taskexecutor.ExecuteTask(taskReq.Command)
	SendTaskResult(taskReq.TaskID, result)
}

func Start() {
	http.HandleFunc("/execute_task", taskHandler)
	log.Println("worker is listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
