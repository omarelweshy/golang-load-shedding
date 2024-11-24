package utils

import (
	"load-shedding/tasks"
	"math/rand"
	"time"
)

func GenerateTask(id int) tasks.Task {
	priorities := []string{"High", "Medium", "Low"}
	duration := time.Duration(rand.Intn(2000)+500) * time.Millisecond

	return tasks.Task{
		ID:       id,
		Priority: priorities[rand.Intn(len(priorities))],
		Duration: duration,
	}
}
