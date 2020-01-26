package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

type key int

const requestIDKey = key(42)

// Println calls log.Printf to print to the standard logger but adds the
// request from the event context.
// START OMIT
func Println(ctx context.Context, msg string) {
	id, ok := ctx.Value(requestIDKey).(int64)
	if !ok {
		log.Printf("[unknown] %s", msg)
		return
	}
	log.Printf("[%d] %s", id, msg)
}

func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, requestIDKey, id)
		f(w, r.WithContext(ctx))
	}
}

// END OMIT
