package game

import (
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	greet      string
	idx        int
	lastLetter time.Time
}

func NewGame(g string) Game {
	return Game{greet: g, idx: 0}
}

func (g *Game) Update() error {
	if g.idx >= len(g.greet) {
		return nil
	}
	now := time.Now()
	if now.Sub(g.lastLetter) >= 500*time.Millisecond {
		g.idx++
		g.lastLetter = now
		log.Printf("%s, counter: %d", now, g.idx)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0, 0, 0xff})
	ebitenutil.DebugPrint(screen, g.greet[:g.idx])
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
