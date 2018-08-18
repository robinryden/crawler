package main

import (
	"time"
)

var JobQueue = make(chan Job)

type Job struct {
	url  string
	time time.Time
}

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func CreateWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

// Start listening on incoming works
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				job.Crawl(job.url)
			case <-w.quit:
				return
			}
		}
	}()
}

// Stop listening on incoming works
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
