package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	c := make(chan int)
	r, w := io.Pipe()
	var buf = make([]byte, 64)
	go func(data []byte, c chan int) {
		_, err := w.Write(data)
		if err != nil {
			fmt.Println(err)
		}
		c <- 0
	}([]byte("hello"), c)
	n, err := r.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf[:n]))
	<-c
	r.Close()
	w.Close()

	PipeSequence()
}

func PipeSequence() {
	c := make(chan int)
	r, w := io.Pipe()
	go func(r io.Reader, c chan int) {
		var buf = make([]byte, 64)
		for {
			n, err := r.Read(buf)
			if err == io.EOF {
				c <- 0
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			c <- n
		}
	}(r, c)
	var buf = make([]byte, 64)
	for i := 0; i < 5; i++ {
		p := buf[0 : 5+i*10]
		n, err := w.Write(p)
		if err != nil {
			log.Fatal(err)
		}
		nn := <-c
		fmt.Printf("Wrote %d, read got %d\n", n, nn)
	}
	w.Close()
	nn := <-c
	fmt.Printf("got %d", nn)
}
