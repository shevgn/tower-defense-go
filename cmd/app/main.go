// Package app provides the main application
package main

import (
	"time"

	"tower-defense-go/pkg/terminal/examples"
)

func main() {
	examples.RunTypewriter(time.Millisecond * 1000)

	examples.RunRectangle(time.Millisecond * 2000)

	examples.RunLines(time.Millisecond * 2000)
}
