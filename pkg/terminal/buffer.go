package terminal

import (
	"strings"
	"sync"
)

// Attribute represents a terminal cell attribute
type Attribute struct {
	Fg Color
	Bg Color
}

// Cell represents a terminal cell
type Cell struct {
	Ch   rune
	Attr Attribute
}

// Buffer represents a terminal buffer
type Buffer struct {
	W, H  int
	Cells []Cell // len = W*H
	mu    sync.RWMutex
}

// NewBuffer creates a new buffer
func NewBuffer(w, h int) *Buffer {
	return &Buffer{
		W:     w,
		H:     h,
		Cells: make([]Cell, w*h),
	}
}

// Index returns the index of the cell at (x,y)
func (b *Buffer) Index(x, y int) int {
	return x + y*b.W
}

// Set sets the cell at (x,y) to cell
func (b *Buffer) Set(x, y int, cell Cell) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if x < 0 || x >= b.W || y < 0 || y >= b.H {
		return
	}
	b.Cells[b.Index(x, y)] = cell
}

// SetCh sets the character at (x,y) to ch
func (b *Buffer) SetCh(x, y int, ch rune) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if x < 0 || x >= b.W || y < 0 || y >= b.H {
		return
	}
	b.Cells[b.Index(x, y)].Ch = ch
}

// Get returns the character at (x,y)
func (b *Buffer) Get(x, y int) *Cell {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if x < 0 || x >= b.W || y < 0 || y >= b.H {
		return nil
	}
	return &b.Cells[b.Index(x, y)]
}

// GetCh returns the character at (x,y)
//
// It is equivalent to calling Get(x,y).Ch
func (b *Buffer) GetCh(x, y int) rune {
	return b.Get(x, y).Ch
}

// Fill fills the buffer with ch
func (b *Buffer) Fill(ch rune) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for i := range b.Cells {
		b.Cells[i].Ch = ch
	}
}

// RowString returns the string representation of the row y (w/o ending '\n')
//
// NOTE: does not take into account the width of wide characters.
func (b *Buffer) RowString(y int) string {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if y < 0 || y >= b.H {
		return ""
	}

	var sb strings.Builder
	start := y * b.W
	for i := 0; i < b.W; i++ {
		sb.WriteRune(b.Cells[start+i].Ch)
	}
	return sb.String()
}

// Copy returns a deep copy of the buffer
func (b *Buffer) Copy() *Buffer {
	b.mu.RLock()
	defer b.mu.RUnlock()

	nb := NewBuffer(b.W, b.H)
	copy(nb.Cells, b.Cells)
	return nb
}
