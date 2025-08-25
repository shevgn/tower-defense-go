// Package app provides the main application
package main

import (
	"fmt"
	"time"

	"tower-defense-go/pkg/terminal"
)

var _ = []rune{
	'#', '#', '#', '#', '#', '#', '#', '#', '#', '#',
	'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#',
	'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#',
	'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#',
	'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#',
	'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#',
	'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#',
	'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#',
	'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#',
	'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#',
	'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#',
	'#', '#', '#', '#', '#', '#', '#', '#', '#', '#',
}

func main() {
	terminal := terminal.NewTerminal()
	terminal.Default()
	terminal.RawMode()
	defer terminal.Restore()
	terminal.Clear()

	msg := "Welcome to the game!"
	terminal.MoveTo(terminal.Width()/2-len(msg)/2, terminal.Height()/2)
	fmt.Print(msg)

	// 	reader := bufio.NewReader(os.Stdin)
	// 	for {
	// 		fmt.Print("> ")
	// 		text, _, _ := reader.ReadRune()
	// 		fmt.Println(text)
	// 	}

	time.Sleep(time.Millisecond * 2000)

	terminal.Clear()
}

type Game struct {
	Width  int
	Height int
}

func NewGame(width, height int) *Game {
	return &Game{
		Width:  width,
		Height: height,
	}
}

func (g *Game) Run() {
	for i := range g.Height {
		for j := range g.Width {
			time.Sleep(time.Millisecond * 1)

			if i == 0 || i == g.Height-1 {
				fmt.Print("#")
				continue
			}
			if j == 0 || j == g.Width-1 {
				fmt.Print("#")
				continue
			}

			fmt.Printf(" ")

			_ = i
			_ = j
		}
	}
}

func (g *Game) Update() {
}
