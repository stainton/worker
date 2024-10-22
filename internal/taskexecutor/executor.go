package taskexecutor

import (
	"log"
	"os/exec"
)

func ExecuteTask(command string) string {
	cmd := exec.Command("/bin/sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Task execution failed: %v", err)
		return "failed"
	}
	log.Printf("Task execute successfully: %s", string(output))
	return "completed"
}
