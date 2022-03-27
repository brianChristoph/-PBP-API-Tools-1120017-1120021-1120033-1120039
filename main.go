package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

func task() {
	fmt.Println("Running a task")
}

func taskWithParams(str string) {
	fmt.Println(str)
}

func CronConcept() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Seconds().Do(task)
	s.Every(1).Second().Do(taskWithParams, "jeda")

	_, err := s.Every(10).Seconds().Do(taskWithParams, "10 detik")
	if err != nil {
		log.Fatalf("Error creating job: %v", err)
	}
	s.StartAsync()
	s.StartBlocking()
}

func main() {
	CronConcept()
	GoRoutine()
	// WithoutGoRoutine()
	
	GoMail()
}
