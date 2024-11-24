package workers

import (
	"fmt"
	"load-shedding/manager"
	"sync"
	"time"
)

func StartWorker(id int, lm *manager.LoadManager, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		task := lm.GetNextTask()
		if task == nil {
			time.Sleep(100 * time.Millisecond) // Wait for new tasks
			continue
		}

		fmt.Printf("Worker %d processing task %d (Priority: %s)\n", id, task.ID, task.Priority)
		time.Sleep(task.Duration)
	}
}
