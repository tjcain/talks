package main

import (
	"fmt"
	"time"
)

func main() {
	bottlesOfBeer(99)
}

// START OMIT
func bottlesOfBeer(n int) {
	for n > 0 {
		fmt.Printf("%[1]d bottles of 🍺 on the wall,\n%[1]d bottles of 🍺,\n", n)
		fmt.Printf("take one down pass it around,\n%d bottles of 🍺 on the wall.\n\n", n-1)
		time.Sleep(2 * time.Second)
		n--
	}
}

// END OMIT
