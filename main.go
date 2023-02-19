package main

import (
	"log"

	"16bchess/chess"
	"16bchess/graphics"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WINDOW_WIDTH  = 1280
	WINDOW_HEIGHT = 960

	PIXEL_WIDTH  = 320
	PIXEL_HEIGHT = 240
)

type Game struct {
	pixels  []byte
	board   chess.Board
	colours graphics.Colours
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.colours.Board)

	g.board.Render(screen, graphics.Position{(PIXEL_WIDTH / 2) - g.board.Size/2, (PIXEL_HEIGHT / 2) - g.board.Size/2})
	//screen.WritePixels(g.pixels)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return PIXEL_WIDTH, PIXEL_HEIGHT
}

func main() {
	game := Game{}
	game.colours = *graphics.NewColours()
	game.board = *chess.NewBoard(100, game.colours)

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Chess")

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
