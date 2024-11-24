package manager

import (
	"fmt"
	"load-shedding/tasks"
	"sync"
)

type LoadManager struct {
	mu          sync.Mutex
	taskQueue   []tasks.Task
	maxCapacity int
}

func NewLoadManager(capacity int) *LoadManager {
	return &LoadManager{
		taskQueue:   make([]tasks.Task, 0),
		maxCapacity: capacity,
	}
}

func (lm *LoadManager) AcceptTask(task tasks.Task) bool {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	if len(lm.taskQueue) >= lm.maxCapacity {
		if task.Priority == "Low" {
			fmt.Printf("Task %d dropped (Low priority)\n", task.ID)
			return false
		}
		lm.shedLoad(task.Priority)
	}

	lm.taskQueue = append(lm.taskQueue, task)
	fmt.Printf("Task %d accepted (Priority: %s)\n", task.ID, task.Priority)
	return true
}

func (lm *LoadManager) shedLoad(minPriority string) {
	priorities := map[string]int{"High": 1, "Medium": 2, "Low": 3}

	newQueue := make([]tasks.Task, 0)
	for _, task := range lm.taskQueue {
		if priorities[task.Priority] < priorities[minPriority] {
			newQueue = append(newQueue, task)
		} else {
			fmt.Printf("Task %d shed (Priority: %s)\n", task.ID, task.Priority)
		}
	}
	lm.taskQueue = newQueue
}

func (lm *LoadManager) GetNextTask() *tasks.Task {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	if len(lm.taskQueue) == 0 {
		return nil
	}

	task := lm.taskQueue[0]
	lm.taskQueue = lm.taskQueue[1:]
	return &task
}
