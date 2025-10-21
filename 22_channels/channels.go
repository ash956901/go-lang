package main

import (
	"fmt"
	"time"
)

// sending
// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("processing number", num)
// 		time.Sleep(time.Second * 1)
// 	}
// }

// receiving
// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult // blocking in sending receiving --> partial truth(buffering channel)
// }

// syncrhonization
// func task(done chan bool) {
// 	defer func() { done <- true }()
// 	fmt.Println("processing...")

// }

// By default channles are unbuffered
// that is sending and receiving are blocking
// to implement something like a queue system
// we can make use of buffered channels
func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()
	// range is infinite blocks this loop, it keeps on going so we need to close the channel
	for email := range emailChan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Second)
	}
}

// channels --> like a pipe from one end
// we send data from another we receive
// helps send data from one go routine to another
// helps in go routine communication
func main() {

	// how to listen to multiple channels
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "ping"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received data from chan2", chan2Val)

		}
	}
	// //buffered channel
	// emailChan := make(chan string, 100) //in real life we send struct instead of string
	// done := make(chan bool)
	// //if after 100 size we send data then those operations become blocking
	// //within the size they remain non blocking

	// go emailSender(emailChan, done)

	// for i := 0; i < 5; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("done sending....")
	// emailChan <- "10@example.com" //now not blocking
	// emailChan <- "20@example.com" //wont go into deadlock

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	// *****************Close the channel*************************
	// close(emailChan)
	// <-done //blocked our goroutine of done which is blocking

	// done := make(chan bool)
	// go task(done)
	// <-done // receive to block (no need to use time or wait group here)

	// result := make(chan int)

	// go sum(result, 5, 4)

	// res := <-result
	// fmt.Println("result:", res)

	// numChan := make(chan int)
	// go processNum(numChan)
	// // numChan <- 5
	// for {
	// 	numChan <- rand.IntN(100)
	// }
	//we utilize channels as sort of queues
	//in go lang

	// messageChan :=make(chan string)
	//kind of data you want to send in channel.

	//sending data into channel
	// messageChan <- "ping" //channels are blocking
	//it is blocked until the second side isnt ready to receive it

	//receive data from channel
	// msg:=<-messageChan

	// fmt.Println(msg)
	// will give error : all goroutines are asleep -->deadlock
}

// Advice : to block single goroutine --> channel
// to block multiple --> wait group
