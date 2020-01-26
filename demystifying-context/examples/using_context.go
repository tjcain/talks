package main

import (
	"context"
	"time"
)

func main() {
	// START0 OMIT
	ctx := context.Background()
	// END0 OMIT

	// START1 OMIT
	ctx, cancel := context.WithCancel(ctx)
	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()
	// END1 OMIT
	// START2 OMIT
	ctx, cancel := context.WithTimeOut(ctx, time.Second)
	defer cancel()
	// END2 OMIT
	// START3 OMIT
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	// END3 OMIT

	// START4 OMIT
	ctx := context.WithValue(ctx, key, id)
	// END4 OMIT

	// START5 OMIT
	rootCtx = context.Background()
	childCtx1 = context.WithValue(rootCtx, key, id)
	childCtx2, cancel := context.WithCancel(rootCtx)
	childCtx3 := context.WithValue(rootCtx, key, id)
	childCtx4 := context.WithValue(childCtx1, key, id)
	// END5 OMIT
}
