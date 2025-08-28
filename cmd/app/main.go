// Package app provides the main application
package main

import (
	"os"
	"time"

	"tower-defense-go/pkg/terminal"
)

const (
	FPS           = 15
	FrameDuration = time.Second / FPS
)

func main() {
	trm := terminal.Default()
	if err := trm.RawMode(); err != nil {
		panic(err)
	}
	defer trm.Restore()
	defer trm.Clear()

	trm.HideCursor()
	defer trm.ShowCursor()

	r := terminal.NewRendererFrom(trm.Width(), trm.Height())
	r.Buf.Fill(' ')
	for x := 0; x < r.Buf.W; x++ {
		r.Buf.Set(x, 0, terminal.BorderHorizontal)
		r.Buf.Set(x, r.Buf.H-1, terminal.BorderHorizontal)
	}
	for y := 0; y < r.Buf.H; y++ {
		r.Buf.Set(0, y, terminal.BorderVertical)
		r.Buf.Set(r.Buf.W-1, y, terminal.BorderVertical)
	}
	r.Buf.Set(0, 0, terminal.BorderRoundedTopLeft)
	r.Buf.Set(r.Buf.W-1, 0, terminal.BorderRoundedTopRight)
	r.Buf.Set(0, r.Buf.H-1, terminal.BorderRoundedBottomLeft)
	r.Buf.Set(r.Buf.W-1, r.Buf.H-1, terminal.BorderRoundedBottomRight)

	_ = r.FlushFull(os.Stdout)

	time.Sleep(time.Second * 2) // Sleep a bit to see the effect
}
