package terminal

import (
	"fmt"
	"log"
	"strings"
)

// Renderer represents a renderer
type Renderer struct {
	cursor *Cursor

	currentColor Color
	resetColor   Color

	termWidth  int
	termHeight int
}

type Rect struct {
	X, Y, Width, Height int
}

type Box struct {
	X, Y, Side int
}

type Line struct {
	X, Y, Length int
}

// NewRenderer creates a new renderer
func NewRenderer() *Renderer {
	return &Renderer{
		cursor:       NewCursor(),
		currentColor: ColorReset,
		resetColor:   ColorReset,
	}
}

// SetSize sets the renderer size
func (r *Renderer) SetSize(width, height int) {
	r.termWidth = width
	r.termHeight = height
}

// SetColor sets the current color
func (r *Renderer) SetColor(color Color) {
	r.currentColor = color
	fmt.Print(r.currentColor)
}

// SetResetColor sets the reset color
func (r *Renderer) SetResetColor(color Color) {
	r.resetColor = color
}

// ResetColor resets the current color
func (r *Renderer) ResetColor() {
	r.currentColor = r.resetColor
	fmt.Print(r.resetColor)
}

// Cursor returns the renderer cursor
func (r *Renderer) Cursor() *Cursor {
	return r.cursor
}

// DrawRect draws a rectangle on the terminal
//
// x and y are the top left corner coordinates.
// width and height are the dimensions of the rectangle.
// The rectangle is filled if the fill parameter is true.
func (r *Renderer) DrawRect(x, y, width, height int, fill bool) {
	if width <= 0 || height <= 0 {
		log.Printf("Width and height must be greater than 0")
		return
	}

	for yi := range height {
		var sb strings.Builder

		for xi := range width {
			if ch, ok := r.borderAt(xi, yi, width, height); ok {
				sb.WriteRune(ch)
				continue
			}

			if fill {
				sb.WriteRune('#')
			} else {
				sb.WriteRune(' ')
			}
		}

		r.cursor.PrintAt(x, y+yi, sb.String())
	}
}

// DrawBox draws a box on the terminal
//
// DrawBox actually calls [terminal.DrawRect] under the hood but with width twice the height
func (r *Renderer) DrawBox(x, y, side int, fill bool) {
	r.DrawRect(x, y, side*2, side, fill)
}

// DrawLineH draws a horizontal line on the terminal
func (r *Renderer) DrawLineH(x, y, length int) {
	if length <= 0 {
		return
	}
	var sb strings.Builder
	for range length {
		sb.WriteRune(BorderHorizontal)
	}
	r.cursor.PrintAt(x, y, sb.String())
}

// DrawLineV draws a vertical line on the terminal
func (r *Renderer) DrawLineV(x, y, length int) {
	r.cursor.MoveTo(x, y)

	for i := range length {
		r.cursor.PrintRuneAt(x, y+i, rune(BorderVertical))
	}
}

func (r *Renderer) borderAt(x, y, width, height int) (rune, bool) {
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

	return ' ', false
}
