package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	ExampleDuration()
	ExampleAfter()
	ExampleSleep()
	ExampleMonth()
	ExampleDate()
	ExampleFormat()
	ExampleParse()
	ExampleParseInLocation()
	//ExampleTick()
	ExampleTickDois()
}

func ExampleDuration() {
	t0 := time.Now()
	time.Sleep(1000)
	t1 := time.Now()
	fmt.Printf("Duration: %v\n", t1.Sub(t0))
}

func ExampleAfter() {
	// select times out because channel does not receive anything
	c := make(chan int)
	select {
	case m := <-c:
		fmt.Println(m)
	case <-time.After(3 * time.Second):
		fmt.Println("timed out")
	}

	//select does not time out because a integer is received on channel c2
	c2 := make(chan int)
	go func(chan int) {
		c2 <- 1
	}(c2)
	select {
	case m2 := <-c2:
		fmt.Println(m2)
	case <-time.After(3 * time.Second):
		fmt.Println("timed out")
	}
}

func ExampleSleep() {
	fmt.Println("Sleeping")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Awakening..")
}

func ExampleTick() {
	//ticks every 2 seconds
	c := time.Tick(2 * time.Second)
	for now := range c {
		fmt.Println(now)
	}
}

func ExampleMonth() {
	year, month, day := time.Now().Date()
	fmt.Println("Year: ", year)
	fmt.Println("Month: ", month)
	fmt.Println("Day: ", day)
}

func ExampleDate() {
	t := time.Date(2014, time.December, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Blalala at %s\n", t.Local())
}

func ExampleFormat() {
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	t := time.Date(2009, time.November, 10, 15, 0, 0, 0, time.Local)
	fmt.Println(t.Format(layout))
	fmt.Println(t.UTC().Format(layout))
}

func ExampleParse() {
	//first parameter is the layout
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println(t)

	//no timezone present
	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2013-Feb-03")
	fmt.Println(t)
}

func ExampleParseInLocation() {
	loc, _ := time.LoadLocation("Europe/Berlin")

	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CEST)", loc)
	fmt.Println(t)
}

func ExampleTickDois() {
	ticker := time.NewTicker(100 * time.Millisecond)
	t0 := time.Now()
	for i := 0; i < 10; i++ {
		<-ticker.C
	}
	ticker.Stop()
	t1 := time.Now()
	fmt.Println(t1.Sub(t0))
	select {
	case <-ticker.C:
		log.Fatal("ticker did not shut down")
	default:
		fmt.Println("ticker did shut down")
	}
}
