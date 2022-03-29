package controllers

import (
	"fmt"
	"time"
)

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func GoRoutine() {
	fmt.Println("Concurrency with GoRoutines")

	go compute(5)
	go compute(5)

	fmt.Scanln()
}

func WithoutGoRoutine() {
	fmt.Println("Concurrency with GoRoutines")

	compute(5)
	compute(5)
}
