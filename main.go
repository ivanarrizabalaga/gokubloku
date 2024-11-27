package main

import (
	"gokudoku/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	g := game.NewGame("Hello gokudoku!!")
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
