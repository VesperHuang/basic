package main

import (
	"fmt"
	"time"
)

// func worker(id int, c chan int) {
// 	for {
// 		// n := <-c
// 		// fmt.Println(n)

// 		fmt.Printf("Worker %d received %c\n", id, <-c)
// 	}
// }

func worker(id int, c chan int) {
	// for {
	// 	n, ok := <-c
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Printf("Worker %d received %c\n", id, n)
	// }

	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

func creatWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	// var c chan int // c == nil
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = creatWorker(i)
	}

	// go func() {
	// 	for {
	// 		n := <-c
	// 		fmt.Println(n)
	// 	}
	// }()
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
	fmt.Println("Buffered channel")
	bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()
}
