package main // OMIT

import ( // OMIT
	"fmt"  // OMIT
	"time" // OMIT
) // OMIT

// START0 OMIT
// The pinger prints a ping and waits for a pong
func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		val := <-pinger
		fmt.Println("ping")
		time.Sleep(500 * time.Millisecond)
		ponger <- val + 1
	}
}

// END0 OMIT

// START1 OMIT
// The ponger prints a pong and waits for a ping
func ponger(pinger chan<- int, ponger <-chan int) {
	for {
		<-ponger
		fmt.Println("pong")
		time.Sleep(200 * time.Millisecond)
		pinger <- 1
	}
}

// END1 OMIT

// START2 OMIT
func main() {
	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)

	ping <- 1 // start ping/ping by sending into the ping channel
	time.Sleep(20 * time.Second)
}

// END2 OMIT
