package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type key int

const (
	beerxIdKey key = iota
	fooIdKey
	barIdKey
)

func main() {

	// START OMIT
	beers := []int{20, 30, 40, 50, 60}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	wg := sync.WaitGroup{}
	for i, b := range beers {
		wg.Add(1)
		ctx := context.WithValue(ctx, idKey, i) // HL
		go bottlesOfBeer(ctx, b, &wg)
	}

	time.Sleep(5 * time.Second) // "proper" blocking exempt for brevity
	cancel()
	wg.Wait()
	// END OMIT
}

func bottlesOfBeer(ctx context.Context, n int, wg *sync.WaitGroup) {
	id, ok := ctx.Value(beerIdKey).(int64)
	if !ok {
		log.Printf("unknown beerIdKey")
	}

	for {
		select {
		case <-ctx.Done():
			log.Println(ctx.Err())
			wg.Done()
			return
		default:
			fmt.Printf("%[1]d bottles of 🍺 on the wall,\n%[1]d bottles of 🍺,\n", n)
			fmt.Printf("take one down pass it around,\n%d bottles of 🍺 on the wall.\n\n", n-1)
			time.Sleep(2 * time.Second)
			n--
			if n <= 0 {
				wg.Done()
				return
			}
		}
	}
}
