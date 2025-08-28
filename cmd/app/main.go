// Package app provides the main application
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"tower-defense-go/pkg/terminal"
)

const (
	FPS           = 60
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	defer func() {
		signal.Stop(sigCh)
		close(sigCh)
	}()

	go func() {
		select {
		case <-sigCh:
			cancel()
		case <-ctx.Done():
		}
	}()

	go func() {
		buf := make([]byte, 1)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil || n == 0 {
				cancel()
				return
			}
			b := buf[0]
			if b == 0x03 { // Ctrl+C
				cancel()
				return
			}
			if b == 0x04 { // Ctrl+D (EOF)
				cancel()
				return
			}
			// handle other keys (arrows, enter, etc.)
		}
	}()

	r := terminal.NewRendererFrom(trm.Width(), trm.Height())

	drawBorder := func(buf *terminal.Buffer) {
		for x := 0; x < buf.W; x++ {
			buf.SetCh(x, 0, terminal.BorderHorizontal)
			buf.SetCh(x, buf.H-1, terminal.BorderHorizontal)
		}
		for y := 0; y < buf.H; y++ {
			buf.SetCh(0, y, terminal.BorderVertical)
			buf.SetCh(buf.W-1, y, terminal.BorderVertical)
		}
		r.Buf.SetCh(0, 0, terminal.BorderRoundedTopLeft)
		r.Buf.SetCh(r.Buf.W-1, 0, terminal.BorderRoundedTopRight)
		r.Buf.SetCh(0, r.Buf.H-1, terminal.BorderRoundedBottomLeft)
		r.Buf.SetCh(r.Buf.W-1, r.Buf.H-1, terminal.BorderRoundedBottomRight)
	}

	ticker := time.NewTicker(FrameDuration)
	defer ticker.Stop()

	i := 0
	for {
		select {
		case <-ctx.Done():
			// final cleanup/animation
			fmt.Fprintln(os.Stderr, "Exiting...")
			return
		case <-ticker.C:
			r.Buf.Fill(' ')
			drawBorder(r.Buf)

			x := i % (r.Buf.W - 2)
			r.Buf.SetCh(1+x, r.Buf.H/2, 'E')

			if err := r.FlushFull(os.Stdout); err != nil {
				fmt.Fprintln(os.Stderr, "Render error:", err)
				cancel()
				return
			}

			i++
		}
	}
}
