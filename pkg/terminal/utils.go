package terminal

import (
	"context"
	"fmt"
	"os"
	"time"
	"unicode/utf8"
)

// CenterX returns the centered x coordinate for the given message
func CenterX(screenWidth int, msg string) int {
	runes := utf8.RuneCountInString(msg)
	x := screenWidth/2 - runes/2
	if x < 0 {
		return 0
	}
	return x
}

func Typewriter(ctx context.Context, w *os.File, msg string, delay time.Duration) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for _, r := range msg {
			select {
			case <-ctx.Done():
				return
			default:
			}

			fmt.Fprintf(w, "%c", r)

			sleep := delay
			if r == ' ' || r == '\n' {
				sleep = delay * 3
			}

			t := time.NewTimer(sleep)
			select {
			case <-ctx.Done():
				t.Stop()
				return
			case <-t.C:
			}
		}
	}()
	return done
}
