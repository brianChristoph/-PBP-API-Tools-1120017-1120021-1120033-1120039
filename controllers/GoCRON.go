package controllers

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

// Begin GoCron
func CronConcept1() {
	// Membuat GoCron
	s := gocron.NewScheduler(time.UTC)

	// Contoh kasus
	// Detik
	s.Every(5).Seconds().Do(task)
	s.Every(1).Second().Do(taskWithParams, "jeda")

	// Menit
	s.Every("5m").Do(task)

	// Hari
	s.Every(5).Days().Do(task)

	// Bulan
	s.Every(2).Month(1, 2, 3).Do(taskWithParams, "bulan") // Januari, Februari, Maret

	// set time tertentu
	s.Every(1).Day().At("10:30").Do(task)

	// set multiple time
	s.Every(1).Day().At("10:30;08:00").Do(taskWithParams, "cara 1 set multiple time")
	s.Every(1).Day().At("10:30").At("08:00").Do(taskWithParams, "cara 2 set multiple time")

	// Schedule each last day of the month
	s.Every(1).MonthLastDay().Do(taskWithParams, "Last day of the month")

	// Or each last day of every other month
	s.Every(2).MonthLastDay().Do(taskWithParams, "Last day of every other month")

	// Misalkan ingin menjalankan suatu fungsi setiap detiknya, dan ingin cek apakah ada error atau tidak
	// Bisa assign 2 variable utk nampung success dan error dari run go cron
	_, err := s.Every(10).Seconds().Do(taskWithParams, "10 detik")
	if err != nil {
		log.Fatalf("Error creating job: %v", err)
	}

	// 2 cara ngerun scheduler
	// Start scheduler secara asynchronously
	s.StartAsync()
	// start scheduler dan block current execution path
	s.StartBlocking()
}

// Cara kedua - Error
func CronConcept2() {
	// Create Go Cron
	s := gocron.NewScheduler(time.UTC)

	/*
		@yearly (or @annually) 	= run once a year, midnight, Jan. 1st 			= 0 0 0 1 1 *
		@montly 				= run once a month, midnight, first of month 	= 0 0 0 1 * *
		@wekkly 				= run once a week, midnight between Sat/Sun 	= 0 0 0 * * 0
		@daily (or @midnight)	= run once a day, midnight 						= 0 0 0 * * *
		@hourly 				= run once an hour, beginning of hour 			= 0 0 * * * *
		@every <duration>
	*/
	s.Cron("@every 1s").Do(taskWithParams, "second")
	s.Cron("*/1 * * * *").Do(taskWithParams, "Every minute")
	s.StartAsync()
	s.StartBlocking()
}
