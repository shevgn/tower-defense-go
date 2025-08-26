package terminal

import "fmt"

// Renderer represents a renderer
type Renderer struct {
	cursor *Cursor

	currentColor Color
	resetColor   Color
}

// NewRenderer creates a new renderer
func NewRenderer() *Renderer {
	return &Renderer{
		cursor:       NewCursor(),
		currentColor: ColorReset,
		resetColor:   ColorReset,
	}
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
	r.cursor.MoveTo(x, y)

	for i := range height {
		for j := range width {

			borderType, ok := r.borderAt(j, i, width, height)
			if ok {
				r.cursor.PrintRuneAt(x+j, y+i, rune(borderType))
				continue
			}

			if fill {
				r.cursor.PrintRuneAt(x+j, y+i, '#')
				continue
			}
		}
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
	r.cursor.MoveTo(x, y)

	for i := range length {
		r.cursor.PrintRuneAt(x+i, y, rune(BorderHorizontal))
	}
}

// DrawLineV draws a vertical line on the terminal
func (r *Renderer) DrawLineV(x, y, length int) {
	r.cursor.MoveTo(x, y)

	for i := range length {
		r.cursor.PrintRuneAt(x, y+i, rune(BorderVertical))
	}
}

func (r *Renderer) borderAt(x, y, width, height int) (borderType, bool) {
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
