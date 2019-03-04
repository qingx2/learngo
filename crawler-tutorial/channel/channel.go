package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("\nchanDemo")
	chanDemo()
	fmt.Println("\nbufferChannel")
	bufferChannel()
	fmt.Println("\nchannelClose")
	channelClose()

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

func bufferChannel() {
	c := make(chan int, 3)
	go worker(0, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Millisecond)
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)

	return c
}

func worker(id int, c chan int) {
	for n := range c {
		// n, ok := <-c;
		// if !ok {
		// 	break
		// }

		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

func chanDemo() {
	var channels [10]chan<- int

	for i := 0; i < 10; i++ {
		for i := 0; i < 10; i++ {
			channels[i] = createWorker(i)
		}
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}
