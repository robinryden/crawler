package main

type Dispatcher struct {
	WorkerPool chan chan Job
	quit       chan bool
}

func CreateDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool}
}

func (d *Dispatcher) Run() {
	// run workers
	for i := 0; i < 10; i++ {
		worker := CreateWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		case <-d.quit:
			return
		}
	}
}

func (d *Dispatcher) Stop() {
	go func() {
		d.quit <- true
	}()
}
