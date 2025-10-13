package main

import (
	"fmt"
	"time"
)

// go routines --> lightweight threads for multithreading and concurrency
// a powerful feature

func task(id int) {
	fmt.Println("doing task", id)
}

// main is also go routine
func main() {
	// this is blocking code , that is it runs sequentially
	// for i :=0 ; i<10; i++ {
	// 	task(i)
	// }

	// lets run it parralley/concurrently using go routine
	// it will run it above a seperate lightweight thread
	// they are very performant
	for i := 0; i <= 10; i++ {
		// go task(i)
		go func (i int){
			fmt.Println(i)
		} (i)
	}
	// wont run rightaway ---> gets scheduled by scheduler
	// main got exited , the pending go routines were terminated
	// lets simulate by sleep
	time.Sleep(time.Second * 2)
}
// we can do cpu intensive tasks , run background workers 