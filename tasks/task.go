package tasks

import "time"

type Task struct {
	ID       int
	Priority string
	Duration time.Duration
}
