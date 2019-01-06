package main

import (
	"time"
)

func init() {
	dispatcher := CreateDispatcher(10)
	dispatcher.Run()
}

func main() {
	urls := []string{"https://www.google.se", "https://www.facebook.com"}

	for _, url := range urls {
		job := Job{
			url:  url,
			time: time.Now(),
		}

		JobQueue <- job
	}
}
