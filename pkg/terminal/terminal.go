// Package terminal provides functions for working with the terminal
package terminal

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

const (
	CtrlC = "\x03"
)

// Terminal represents a terminal
type Terminal struct {
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

// Size returns the terminal size
func (t *Terminal) Size() (int, int) {
	return t.width, t.height
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
