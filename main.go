package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	greet string
	idx   int
}

func (g *Game) Update() error {
	if g.idx >= len(g.greet) {
		return nil
	}
	if time.Now().Second()%5 == 0 {
		g.idx++
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.greet[:g.idx])
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	g := Game{
		greet: "Hello world!",
		idx:   0,
	}
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
