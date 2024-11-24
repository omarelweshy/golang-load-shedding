package main

import (
	"load-shedding/manager"
	"load-shedding/utils"
	"load-shedding/workers"
	"sync"
	"time"
)

func main() {
	maxCapacity := 50
	numWorkers := 3
	totalTasks := 50

	loadManager := manager.NewLoadManager(maxCapacity)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go workers.StartWorker(i, loadManager, &wg)
	}

	for id := 1; id <= totalTasks; id++ {
		task := utils.GenerateTask(id)
		loadManager.AcceptTask(task)
		time.Sleep(time.Duration(id%500) * time.Millisecond)
	}

	wg.Wait()
}
