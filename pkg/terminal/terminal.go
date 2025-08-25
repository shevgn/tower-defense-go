// Package terminal provides functions for working with the terminal
package terminal

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

// Terminal represents a terminal
type Terminal struct {
	posX int
	posY int

	state *term.State

	width  int
	height int
}

// NewTerminal creates a new terminal
func NewTerminal() *Terminal {
	return &Terminal{}
}

// Width returns the terminal width
func (t *Terminal) Width() int {
	return t.width
}

// Height returns the terminal height
func (t *Terminal) Height() int {
	return t.height
}

// Default sets the terminal width and height from os.Stdout.Fd()
func (t *Terminal) Default() {
	fd := int(os.Stdout.Fd())

	t.FromFD(fd)
}

// FromFD gets the terminal size from the file descriptor
func (t *Terminal) FromFD(fd int) {
	if !term.IsTerminal(fd) {
		println("Not a terminal")
		os.Exit(1)
	}

	width, height, err := term.GetSize(fd)
	if err != nil {
		panic("Error getting terminal size")
	}

	t.width = width
	t.height = height
}

// RawMode enables raw mode for the terminal
func (t *Terminal) RawMode() {
	fd := int(os.Stdout.Fd())

	prevState, err := term.MakeRaw(fd)
	if err != nil {
		panic("Error getting terminal state")
	}
	t.state = prevState
}

// Restore restores the terminal to its previous state
func (t *Terminal) Restore() {
	fd := int(os.Stdout.Fd())
	_ = term.Restore(fd, t.state)
}

// Clear clears the terminal screen
func (t *Terminal) Clear() {
	fmt.Print("\033[H\033[2J")
}

// MoveTo moves the cursor to the specified position
func (t *Terminal) MoveTo(x, y int) {
	t.posX = x - 1
	t.posY = y - 1
	fmt.Printf("\x1B[%d;%dH", t.posY, t.posX)
}
