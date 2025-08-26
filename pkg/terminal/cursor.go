package terminal

import (
	"fmt"
	"sync"
	"unicode/utf8"
)

// Cursor stores 0-based coordinates internally.
type Cursor struct {
	mu sync.Mutex
	x  int
	y  int
}

// NewCursor creates cursor at (0,0).
func NewCursor() *Cursor {
	return &Cursor{}
}

// MoveTo moves cursor to 0-based (x,y) and emits escape sequence.
func (c *Cursor) MoveTo(x, y int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}
	c.x = x
	c.y = y
	fmt.Printf("\x1B[%d;%dH", c.y+1, c.x+1)
}

// Print prints message at current cursor position and advances cursor.
func (c *Cursor) Print(msg string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	fmt.Printf("\x1b[%d;%dH", c.y+1, c.x+1)
	fmt.Print(msg)
	c.x += utf8.RuneCountInString(msg)
}

// PrintAt prints message at (x,y) (0-based) - optimized: single MoveTo then print.
func (c *Cursor) PrintAt(x, y int, msg string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}
	c.x = x
	c.y = y
	fmt.Printf("\x1b[%d;%dH", c.y+1, c.x+1)
	fmt.Print(msg)
	c.x += utf8.RuneCountInString(msg)
}

// PrintRuneAt prints single rune at (x,y) and advances cursor by 1.
func (c *Cursor) PrintRuneAt(x, y int, r rune) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}
	c.x = x
	c.y = y
	fmt.Printf("\x1b[%d;%dH", c.y+1, c.x+1)
	fmt.Printf("%c", r)
	c.x++
}

// Reset resets the cursor to the top left corner
//
// It is equivalent to calling MoveTo(0, 0)
func (c *Cursor) Reset() {
	c.MoveTo(0, 0)
}

// TopLeft moves cursor to top left corner
//
// It is equivalent to calling MoveTo(0, 0)
func (c *Cursor) TopLeft() { c.MoveTo(0, 0) }

// TopRight moves cursor to top right corner
func (c *Cursor) TopRight() { /* to implement we need terminal width */ }

// BottomRight moves cursor to bottom right corner
func (c *Cursor) BottomRight() { /* to implement we need terminal width */ }

// BottomLeft moves cursor to bottom left corner
func (c *Cursor) BottomLeft() { /* to implement we need terminal height */ }
