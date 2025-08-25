// Package terminal provides functions for working with the terminal
package terminal

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

// Terminal represents a terminal
type Terminal struct {
	state *term.State

	width  int
	height int

	cursor *Cursor
}

// Default sets the terminal width and height from os.Stdout.Fd()
func Default() *Terminal {
	fd := int(os.Stdout.Fd())

	if !term.IsTerminal(fd) {
		println("Not a terminal")
		os.Exit(1)
	}

	width, height, err := term.GetSize(fd)
	if err != nil {
		panic("Error getting terminal size")
	}

	terminal := Terminal{
		width:  width,
		height: height,
		cursor: NewCursor(),
	}

	return &terminal
}

// Width returns the terminal width
func (t *Terminal) Width() int {
	return t.width
}

// Height returns the terminal height
func (t *Terminal) Height() int {
	return t.height
}

// Size returns the terminal size
func (t *Terminal) Size() (int, int) {
	return t.width, t.height
}

// Cursor returns the terminal cursor
func (t *Terminal) Cursor() *Cursor {
	return t.cursor
}

// RawMode enables raw mode for the terminal
func (t *Terminal) RawMode() {
	prevState, err := term.MakeRaw(int(os.Stdout.Fd()))
	if err != nil {
		panic("Error getting terminal state")
	}
	t.state = prevState
}

// Restore restores the terminal to its previous state
func (t *Terminal) Restore() {
	_ = term.Restore(int(os.Stdout.Fd()), t.state)
}

// Clear clears the terminal screen
func (t *Terminal) Clear() {
	fmt.Print("\033[H\033[2J")
}

// DrawRect draws a rectangle on the terminal
//
// x and y are the top left corner coordinates.
// width and height are the dimensions of the rectangle.
// The rectangle is filled if the fill parameter is true.
func (t *Terminal) DrawRect(x, y, width, height int, fill bool) {
	t.cursor.MoveTo(x, y)

	for i := range height {
		for j := range width {

			borderType, ok := t.borderAt(j, i, width, height)
			if ok {
				t.cursor.PrintAt(x+j, y+i, string(borderType))
				continue
			}

			if fill {
				t.cursor.PrintAt(x+j, y+i, "#")
				continue
			}
		}
	}
}

func (t *Terminal) borderAt(x, y, width, height int) (BorderType, bool) {
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
	if y == 0 {
		return BorderTop, true
	}
	// Bottom
	if y == height-1 {
		return BorderBottom, true
	}
	// Left
	if x == 0 {
		return BorderLeft, true
	}
	// Right
	if x == width-1 {
		return BorderRight, true
	}

	return "", false
}
