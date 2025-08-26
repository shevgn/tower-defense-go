// Package app provides the main application
package main

import (
	"fmt"
	"time"

	"tower-defense-go/pkg/terminal"
)

func Typewriter(msg string, delay time.Duration) {
	for _, char := range msg {
		fmt.Printf("%c", char)

		if char == ' ' || char == '\n' {
			time.Sleep(delay * 3)
			continue
		}

		time.Sleep(delay)
	}
}

func main() {
	trm := terminal.Default()
	trm.RawMode()
	defer trm.Restore()
	trm.Clear()

	renderer := trm.Renderer()
	cursor := trm.Cursor()

	msg := "Welcome to the game!"
	cursor.MoveTo(trm.Width()/2-len(msg)/2, trm.Height()/2)
	Typewriter(msg, time.Millisecond*75)
	time.Sleep(time.Millisecond * 1000)
	trm.Clear()

	renderer.DrawRect(1, 1, 20, 20, false)
	time.Sleep(time.Millisecond * 1000)
	trm.Clear()

	renderer.SetColor(terminal.ColorGreen)
	renderer.DrawRect(1, 1, 60, 30, false)
	renderer.ResetColor()
	time.Sleep(time.Millisecond * 1000)
	trm.Clear()

	renderer.SetColor(terminal.ColorRed)
	renderer.DrawBox(1, 1, 20, false)
	renderer.ResetColor()
	time.Sleep(time.Millisecond * 1000)
	trm.Clear()

	renderer.DrawLineH(1, 15, 10)
	renderer.DrawLineV(30, 15, 10)
	renderer.DrawLineV(31, 15, 10)
	renderer.DrawLineV(32, 15, 10)
	renderer.DrawLineV(33, 15, 10)
	renderer.DrawLineV(34, 15, 10)
	renderer.SetColor(terminal.ColorYellow)
	renderer.DrawLineH(20, 20, 20)
	renderer.ResetColor()

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

	time.Sleep(time.Millisecond * 2000)

	trm.Clear()
}
