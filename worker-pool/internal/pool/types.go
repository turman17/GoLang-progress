package pool

import "time"

type Options struct {
	Workers int
	NTasks int
	Sleep int
}

type Result struct {
	JobID int
	Input int
	Output int
	time time.Duration
	Err error
}

type Job struct {
	ID int
	Value int
}