package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done() //--> runs when the function completes
	// like the useEffect cleanup functions return(){}
	fmt.Println("Doing task", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()

	// time.Sleep(time.Second * 2)
	// in real code, we dont know when will these finish
	// so we use wait groups to synchronise them and exit only
	// when all of them got termianated

	// Add() --> +1 for each routine
	// defer --> -1 for each completed routine
	// Wait() --> waits until counter becomes 0
	// wait group --> counter mechanism
}
