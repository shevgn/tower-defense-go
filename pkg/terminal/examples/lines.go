package examples

import (
	"time"

	"tower-defense-go/pkg/terminal"
)

// RunLines runs the lines example
func RunLines(showtime time.Duration) {
	trm := terminal.Default()
	if err := trm.RawMode(); err != nil {
		panic(err)
	}
	defer trm.Restore()
	trm.Clear()

	r := trm.Renderer()

	r.DrawLineH(1, 15, 10)
	r.DrawLineV(20, 15, 10)
	r.SetColor(terminal.BgBlue)
	r.DrawLineV(21, 15, 10)
	r.ResetColor()
	r.DrawLineV(22, 15, 10)
	r.DrawLineV(23, 15, 10)
	r.DrawLineV(24, 15, 10)
	r.SetColor(terminal.FgYellow)
	r.DrawLineH(20, 20, 20)
	r.ResetColor()

	time.Sleep(showtime)
	trm.Clear()
}
