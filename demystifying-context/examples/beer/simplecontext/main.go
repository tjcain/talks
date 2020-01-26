package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	bottlesOfBeer(ctx, 99)
}

// START OMIT
func bottlesOfBeer(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done():
			log.Println(ctx.Err())
			return
		default:
			fmt.Printf("%[1]d bottles of 🍺 on the wall,\n%[1]d bottles of 🍺,\n", n)
			fmt.Printf("take one down pass it around,\n%d bottles of 🍺 on the wall.\n\n", n-1)
			time.Sleep(2 * time.Second)
			n--
			if n <= 0 {
				return
			}
		}
	}
}

// END OMIT
