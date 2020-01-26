package main

import (
	"context"
	"log"
	"sync"
	"time"
)

// START1 OMIT
// do not use basic type to avoid key collisions, pkg1.key(0) != pkg2.key(0)
type myKey int // HL

const (
	idKey myKey = iota
	fooKey
	barKey
)

// END1 OMIT

func main() {
	beers := []int{20, 30, 40, 50, 60}

	// START2 OMIT
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	wg := sync.WaitGroup{}
	for i, b := range beers {
		wg.Add(1)
		ctx := context.WithValue(ctx, idKey, i) // HL
		go bottlesOfBeer(ctx, b, &wg)
	}

	time.Sleep(60 * time.Second) // "proper" blocking exempt for brevity
	cancel()
	wg.Wait()
	// END2 OMIT
}

func bottlesOfBeer(ctx context.Context, n int, wg *sync.WaitGroup) {
	// START3 OMIT
	id, ok := ctx.Value(idKey).(int)
	if !ok {
		log.Print("unkown context key")
		wg.Done() // OMIT
		return    // OMIT
	}
	// END3 OMIT

	for {
		select {
		case <-ctx.Done():
			log.Println(ctx.Err())
			wg.Done()
			return
		default:
			log.Printf("[%d] has %d bottles of 🍺", id, n)
			//fmt.Printf("%[1]d bottles of 🍺 on the wall,\n%[1]d bottles of 🍺,\n", n)
			//fmt.Printf("take one down pass it around,\n%d bottles of 🍺 on the wall.\n\n", n-1)
			time.Sleep(2 * time.Second)
			n--
			if n <= 0 {
				wg.Done()
				return
			}
		}
	}
}
