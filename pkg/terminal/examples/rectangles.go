package examples

import (
	"fmt"
	"time"

	"tower-defense-go/pkg/terminal"
)

// RunRectangle runs the rectangle example
func RunRectangle(showtime time.Duration) {
	trm := terminal.Default()
	if err := trm.RawMode(); err != nil {
		panic(err)
	}
	defer trm.Restore()
	trm.Clear()

	r := trm.Renderer()

	rect := terminal.Rect{X: 0, Y: 0, Width: 30, Height: 20}
	examples := []struct {
		name string
		fn   func()
	}{
		{
			name: fmt.Sprintf("Rectangle %dx%d at (%d,%d)", rect.Width, rect.Height, rect.X, rect.Y),
			fn: func() {
				r.DrawRect(rect.X, rect.Y, rect.Width, rect.Height, false)
			},
		},
		{
			name: fmt.Sprintf("Filled rectangle %dx%d at (%d,%d)", rect.Width, rect.Height, rect.X, rect.Y),
			fn: func() {
				r.DrawRect(rect.X, rect.Y, rect.Width, rect.Height, true)
			},
		},
		{
			name: fmt.Sprintf("Red rectangle %dx%d at (%d,%d)", rect.Width, rect.Height, rect.X, rect.Y),
			fn: func() {
				r.SetColor(terminal.FgRed)
				r.DrawRect(rect.X, rect.Y, rect.Width, rect.Height, false)
				r.ResetColor()
			},
		},
		{
			name: fmt.Sprintf("Blue rectangle %dx%d at (%d,%d)", rect.Width, rect.Height, rect.X, rect.Y),
			fn: func() {
				r.SetColor(terminal.FgBlue)
				r.DrawRect(rect.X, rect.Y, rect.Width, rect.Height, false)
				r.ResetColor()
			},
		},
	}

	for _, e := range examples {
		e.fn()
		r.Cursor().PrintAt(terminal.CenterX(rect.Width, e.name[:len(e.name)/2]), rect.Height/2, e.name[:len(e.name)/2])
		r.Cursor().PrintAt(terminal.CenterX(rect.Width, e.name[len(e.name)/2:]), rect.Height/2+1, e.name[len(e.name)/2:])
		time.Sleep(showtime)
		trm.Clear()
	}

	// b := make([]byte, 1)
	// for {
	// 	n, err := os.Stdin.Read(b)
	// 	if err != nil {
	// 		log.Fatalf("Error reading from stdin: %v", err)
	// 	}
	// 	if n > 0 {
	// 		char := string(b[0])
	// 		fmt.Printf("You pressed: %s\n", char)
	// 		if char == "\x03" {
	// 			break
	// 		}
	// 	}
	// }
}
