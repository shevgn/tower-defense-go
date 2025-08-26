// Package terminal provides functions for working with the terminal
package terminal

import (
	"fmt"
	"os"
	"sync"

	"golang.org/x/term"
)

// Terminal represents a terminal
type Terminal struct {
	state *term.State

	width  int
	height int

	renderer *Renderer

	// fd used for MakeRaw/Restore (stdin)
	fd int
	mu sync.Mutex
}

// Default sets the terminal width and height from os.Stdout.Fd()
func Default() *Terminal {
	inFd := int(os.Stdin.Fd())
	outFd := int(os.Stdout.Fd())

	if !term.IsTerminal(inFd) || !term.IsTerminal(outFd) {
		fmt.Fprintln(os.Stderr, "Not a terminal")
		os.Exit(1)
	}

	width, height, err := term.GetSize(outFd)
	if err != nil {
		panic("Error getting terminal size")
	}

	terminal := Terminal{
		width:    width,
		height:   height,
		renderer: NewRenderer(),
	}

	terminal.renderer.SetSize(width, height)

	return &terminal
}

// Width returns the terminal width
func (t *Terminal) Width() int { return t.width }

// Height returns the terminal height
func (t *Terminal) Height() int { return t.height }

// Size returns the terminal size
func (t *Terminal) Size() (int, int) { return t.width, t.height }

// Cursor returns the terminal cursor
func (t *Terminal) Cursor() *Cursor { return t.renderer.Cursor() }

// Renderer returns the terminal renderer
func (t *Terminal) Renderer() *Renderer { return t.renderer }

// RawMode enables raw mode for the terminal
func (t *Terminal) RawMode() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.fd == 0 {
		t.fd = int(os.Stdin.Fd())
	}
	prevState, err := term.MakeRaw(t.fd)
	if err != nil {
		return err
	}

	t.state = prevState
	return nil
}

// Restore restores the terminal to its previous state
func (t *Terminal) Restore() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.state == nil {
		return nil
	}
	err := term.Restore(t.fd, t.state)
	t.state = nil
	return err
}

// Clear clears the terminal screen
func (t *Terminal) Clear() {
	// fmt.Print("\033[H\033[2J")

	// CSI 2 J = clear, CSI H move to home
	fmt.Print("\x1b[2J\x1b[H")
}
