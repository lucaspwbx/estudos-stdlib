package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// sync with the use of an unbuffered1 channel
	quit1 := make(chan bool)
	go unbuffered1(quit1)
	<-quit1
	fmt.Println("Work Done")

	//using a buffered1 channel and receiving msgs using range
	msgs := make(chan string, 2)
	go buffered1(msgs)
	for msg := range msgs {
		fmt.Println(msg)
	}

	//buffered channel, receiving msgs using for
	msgs2 := make(chan string, 10)
	go buffered2(msgs2)
	for {
		msg, more := <-msgs2
		if more {
			fmt.Println("Received msg ", msg)
		} else {
			fmt.Println("received all msgs")
			break
		}
	}

	msgs3 := make(chan string, 5)
	msgs4 := make(chan string, 10)
	finished := make(chan bool)
	go crazyfunc(msgs3, msgs4, finished)
	for {
		select {
		case foo := <-msgs3:
			fmt.Println(foo)
		case bar := <-msgs4:
			fmt.Println(bar)
		case <-finished:
			fmt.Println("finished")
			return
		}
	}
}

func unbuffered1(quit chan bool) {
	time.Sleep(2 * time.Second)
	quit <- true
}

func buffered1(msgs chan string) {
	msgs <- "hello"
	time.Sleep(1 * time.Second)
	msgs <- "world"
	close(msgs)
}

func buffered2(msgs chan string) {
	for i := 0; i < 10; i++ {
		msgs <- fmt.Sprintf("Message %d", i)
	}
	close(msgs)
}

func crazyfunc(msgs, msgs2 chan string, finished chan bool) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func(msgs chan string) {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			msgs <- fmt.Sprintf("Alou %d", i)
		}
	}(msgs)
	go func(msgs chan string) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			msgs <- fmt.Sprintf("Fuck %d", i)
		}
	}(msgs2)
	wg.Wait()
	finished <- true
}
