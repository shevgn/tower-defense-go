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

// Typewriter prints the given message with a typewriter effect
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

// BorderAt returns the border character at (x,y)
func BorderAt(x, y, width, height int) (rune, bool) {
	// TopLeft
	if x == 0 && y == 0 {
		return BorderRoundedTopLeft, true
	}
	// TopRight
	if y == 0 && x == width-1 {
		return BorderRoundedTopRight, true
	}
	// BottomLeft
	if y == height-1 && x == 0 {
		return BorderRoundedBottomLeft, true
	}
	// BottomRight
	if y == height-1 && x == width-1 {
		return BorderRoundedBottomRight, true
	}
	// Top
	if y == 0 || y == height-1 {
		return BorderHorizontal, true
	}
	// Left
	if x == 0 || x == width-1 {
		return BorderVertical, true
	}

	return ' ', false
}
