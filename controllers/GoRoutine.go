package controllers

import (
	"fmt"
	"time"
)

// Go Routine merupakan function dan method yang berjalan berasamaan dengan function dan method yang lain
// Cost membuat sebuah goroutine jauh lebih ringan dibandingkan dengan membuat sebuah thread
func compute(value int, start int) {
	for i := start; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

//Contoh funtion yang menggunakan go routine
func GoRoutine() {
	fmt.Println("Concurrency WITH GoRoutines")

	go compute(5, 1)
	go compute(6, 2)

	fmt.Scanln()
}

// Contoh function yang tidak menggunakan go routine
func WithoutGoRoutine() {
	fmt.Println("Concurrency WITHOUT GoRoutines")

	compute(5, 1)
	compute(6, 2)
}
