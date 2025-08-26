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

	// cursor.PrintAt(1, 1, "Hello")
	renderer.DrawRect(1, 1, 4, 4, false)

	time.Sleep(time.Millisecond * 1000)

	trm.Clear()

	renderer.DrawBox(1, 1, 4, false)

	renderer.DrawLineH(1, 10, 4)

	renderer.DrawLineV(30, 10, 4)
	renderer.DrawLineV(31, 10, 4)
	renderer.DrawLineV(32, 10, 4)

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
