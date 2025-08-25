package terminal

import "fmt"

// Cursor represents a cursor
type Cursor struct {
	x int
	y int
}

// NewCursor creates a new cursor
func NewCursor() *Cursor {
	return &Cursor{}
}

// MoveTo moves the cursor to the specified position
func (t *Cursor) MoveTo(x, y int) {
	t.x = x + 1
	t.y = y + 1
	fmt.Printf("\x1B[%d;%dH", t.y, t.x)
}

// Print prints a message at the current cursor position
func (t *Cursor) Print(msg string) {
	t.MoveTo(t.x, t.y)

	fmt.Print(msg)

	t.MoveTo(t.x+len(msg), t.y)
}

// PrintAt prints a message at the specified position
func (t *Cursor) PrintAt(x, y int, msg string) {
	t.MoveTo(x, y)

	fmt.Print(msg)
}

// Reset resets the cursor to the top left corner
//
// It is equivalent to calling MoveTo(0, 0)
func (t *Cursor) Reset() {
	t.MoveTo(0, 0)
}

// TopLeft moves the cursor to the top left corner
func (t *Cursor) TopLeft() {
	t.MoveTo(0, 0)
}

// TopRight moves the cursor to the top right corner
func (t *Cursor) TopRight() {
	t.MoveTo(t.x, 0)
}

// BottomRight moves the cursor to the bottom right corner
func (t *Cursor) BottomRight() {
	t.MoveTo(t.x, t.y)
}

// BottomLeft moves the cursor to the bottom left corner
func (t *Cursor) BottomLeft() {
	t.MoveTo(0, t.y)
}
