// Package examples provides examples for the terminal package
package examples

import (
	"context"
	"os"
	"time"

	"tower-defense-go/pkg/terminal"
)

// RunTypewriter runs the typewriter example
func RunTypewriter(showtime time.Duration) {
	trm := terminal.Default()
	if err := trm.RawMode(); err != nil {
		panic(err)
	}
	defer trm.Restore()
	trm.Clear()

	c := trm.Cursor()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	msg := "Welcome to the game!"
	c.MoveTo(terminal.CenterX(trm.Width(), msg), trm.Height()/2)
	<-terminal.Typewriter(ctx, os.Stdout, msg, time.Millisecond*75)
	time.Sleep(showtime)
	trm.Clear()
}
